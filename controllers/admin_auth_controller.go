package controllers

import (
	"apcore/messages"
	"apcore/models"
	"apcore/response"
	"apcore/services"
	"apcore/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AdminSigninMessage struct {
	Token string `json:"token"`
}

type AdminAuthBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AdminAuthController struct {
	service services.AdminService
	jwtService *jwt.JWTService
}

func NewAdminAuthController(service services.AdminService, jwtService *jwt.JWTService) *AdminAuthController {
	return &AdminAuthController{service,jwtService}
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

	admin, err := ctrl.service.GetAdminByName(input.Email)
	if err != nil {
		response.Error(c, nil, messages.MsgInvalidEmailPassword, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		response.Error(c, nil, messages.MsgInvalidEmailPassword, http.StatusUnauthorized)
		return
	}

	token, err := ctrl.jwtService.GenerateJWT(admin.Phone)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)

		return
	}

	response.Success(c, gin.H{"token": token}, messages.MsgSuccessful, nil, http.StatusOK)
}
