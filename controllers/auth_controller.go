package controllers

import (
	"apcore/messages"
	"apcore/models"
	"apcore/response"
	"apcore/services"
	"apcore/utils/jwt"
	"apcore/utils/otp"
	"apcore/utils/sms"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

// @Summary Signup route
// @Description Creates new users
// @Tags auth
// @Accept  json
// @Produce  json
// @Param locale header string true "Locale" Enums(en, fa)
// @Param user body SignupBody true "User Information"
// @Success 201 {object} response.SwaggerResponse[models.User]
// @Router /auth/signup [get]
func (ctrl *AuthController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Error(c, nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	if err := ctrl.service.CreateUser(&user); err != nil {
		response.Error(c, nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Success(c, user, "success", nil, http.StatusCreated)
}

type SigninMessage struct {
	Token string `json:"token"`
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
// @Success 200 {object} response.SwaggerResponse[SigninMessage]
// @Router /auth [post]
func (ctrl *AuthController) Auth(c *gin.Context) {
	var input AuthBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	otp, err := ctrl.otpService.Generate(c.Request.Context(), input.Phone)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	err = ctrl.smsSender.SendLoginOtp(otp, input.Phone)
	if err != nil {
		response.Error(c, nil, "Failed to send OTP", http.StatusInternalServerError)
		return
	}

	response.Success(c, gin.H{"message": "OTP sent successfully"}, messages.MsgSuccessful, nil, http.StatusOK)
}

type VerifySignInBody struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

func (ctrl *AuthController) VerifySignIn(c *gin.Context) {
	var input VerifySignInBody

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

	token, err := ctrl.jwtService.GenerateJWT(user.Phone)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, gin.H{"token": token}, messages.MsgSuccessful, nil, http.StatusOK)
}
