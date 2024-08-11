package controllers

import (
	"apcore/messages"
	"apcore/models"
	"apcore/response"
	"apcore/services"
	"apcore/utils/jwt"
	"apcore/utils/otp"
	"apcore/utils/sms"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupBody struct {
	Username string        `gorm:"unique;not null" json:"username" binding:"required"`
	Email    string        `gorm:"unique;not null" json:"email" binding:"required"`
	Password string        `gorm:"not null" json:"password" binding:"required"`
	Roles    []models.Role `gorm:"many2many:user_roles;" json:"roles" binding:"required"`
}

type AuthController struct {
	service    services.UserService
	jwtService *jwt.JWTService
	otpService *otp.OTPService
	smsSender  sms.SmsSender
}

func NewAuthController(service services.UserService, jwtService *jwt.JWTService, otpService *otp.OTPService, smsSender sms.SmsSender) *AuthController {
	return &AuthController{service, jwtService, otpService, smsSender}
}

type AuthMessage struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthBody struct {
	Phone string `json:"phone" binding:"required"`
}

// @Summary Auth route
// @Description Handles user signin and signup
// @Tags auth
// @Accept  json
// @Produce  json
// @Param locale header string true "Locale" Enums(en, fa)
// @Param user body AuthBody true "User Phone"
// @Success 200 {object} response.SwaggerResponse[AuthMessage]
// @Router /auth [post]
func (ctrl *AuthController) Auth(c *gin.Context) {
	var input AuthBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ctrl.service.GetUserByPhone(input.Phone)
	if err != nil {

		newUser := &models.User{
			FullName: "کاربر گرامی",
			Phone:    input.Phone,
			Nid:      nil,
		}

		if err := ctrl.service.CreateUser(newUser); err != nil {
			response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
			return
		}

		user = newUser
	}

	otp, err := ctrl.otpService.Generate(c.Request.Context(), user.Phone)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	err = ctrl.smsSender.SendLoginOtp(otp, user.Phone)
	if err != nil {
		response.Error(c, nil, "Failed to send OTP", http.StatusInternalServerError)
		fmt.Println("err is: ", err)
		return
	}

	response.Success(c, gin.H{"message": "OTP sent successfully"}, messages.MsgSuccessful, nil, http.StatusOK)
}

type VerifyAuthBody struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

// @Summary Verify Auth route
// @Description Handles user verification
// @Tags auth
// @Accept  json
// @Produce  json
// @Param locale header string true "Locale" Enums(en, fa)
// @Param user body VerifyAuthBody true "Verify OTP"
// @Success 200 {object} response.SwaggerResponse[AuthMessage]
// @Router /auth/verify [post]
func (ctrl *AuthController) VerifyAuth(c *gin.Context) {
	var input VerifyAuthBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	isValid, err := ctrl.otpService.Verify(c, input.Phone, input.OTP)
	if !isValid || err != nil {
		response.Error(c, nil, messages.MsgInvalidOTP, http.StatusUnauthorized)
		return
	}

	user, err := ctrl.service.GetUserByPhone(input.Phone)
	if err != nil {
		response.Error(c, nil, messages.MsgUserNotFound, http.StatusUnauthorized)
		return
	}

	auth, err := ctrl.jwtService.GenerateToken(user.Phone)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, gin.H{"accessToken": auth.AccessToken, "refreshToken": auth.RefreshToken}, messages.MsgSuccessful, nil, http.StatusOK)
}
