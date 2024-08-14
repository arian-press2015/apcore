package controllers

import (
	"apcore/messages"
	"apcore/models"
	"apcore/response"
	"apcore/services"
	"apcore/utils/jwt"
	"apcore/utils/mfa"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminAuthBody struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required"`
	TOTP     string `json:"totp" binding:"required"`
}

type AdminAuthController struct {
	service    services.AdminService
	jwtService *jwt.JWTService
}

func NewAdminAuthController(service services.AdminService, jwtService *jwt.JWTService) *AdminAuthController {
	return &AdminAuthController{service, jwtService}
}

// @Summary Admin auth route
// @Description Logs in the admin
// @Tags admin
// @Accept  json
// @Produce  json
// @Param locale header string true "Locale" Enums(en, fa)
// @Param user body AdminAuthBody true "Admin Credentials"
// @Success 200 {object} response.SwaggerResponse[AuthMessage]
// @Router /admin/auth [post]
func (ctrl *AdminAuthController) AdminLogin(c *gin.Context) {
	var admin *models.Admin
	var input AdminAuthBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	admin, err := ctrl.service.GetAdminByPhone(input.Phone)
	if err != nil {
		response.Error(c, nil, messages.MsgInvalidPhonePassword, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		response.Error(c, nil, messages.MsgInvalidPhonePassword, http.StatusUnauthorized)
		return
	}

	if !mfa.VerifyTOTPCode(admin.TOTPSecret, input.TOTP) {
		response.Error(c, nil, messages.MsgUnauthorized, http.StatusUnauthorized)
		return
	}

	auth, err := ctrl.jwtService.GenerateToken(admin.Phone)
	if err != nil {
		response.Error(c, nil, "Failed to generate JWT", http.StatusInternalServerError)
		return
	}

	response.Success(c, gin.H{"accessToken": auth.AccessToken, "refreshToken": auth.RefreshToken}, messages.MsgSuccessful, nil, http.StatusOK)
}

// func (ctrl *AdminAuthController) CreateAdmin(c *gin.Context) {

// 	secret, qrCodeURL, err := mfa.GenerateTOTPSecret(admin.Phone)
// 	if err != nil {
// 		response.Error(c, nil, "Failed to generate TOTP secret", http.StatusInternalServerError)
// 		return
// 	}

// 	admin.TOTPSecret = secret
// }
