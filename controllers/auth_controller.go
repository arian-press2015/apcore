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

type SignupBody struct {
	Username string        `gorm:"unique;not null" json:"username" binding:"required"`
	Email    string        `gorm:"unique;not null" json:"email" binding:"required"`
	Password string        `gorm:"not null" json:"password" binding:"required"`
	Roles    []models.Role `gorm:"many2many:user_roles;" json:"roles" binding:"required"`
}

type AuthController struct {
	service services.UserService
	jwtService *jwt.JWTService
}

func NewAuthController(service services.UserService, jwtService *jwt.JWTService) *AuthController {
	return &AuthController{service, jwtService}
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

type SigninBody struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Signin route
// @Description Logs in the user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param locale header string true "Locale" Enums(en, fa)
// @Param user body SigninBody true "User Credentials"
// @Success 200 {object} response.SwaggerResponse[SigninMessage]
// @Router /auth/signin [get]
func (ctrl *AuthController) Login(c *gin.Context) {
	var user *models.User
	var input SigninBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := ctrl.service.GetUserByUsername(input.Email)
	if err != nil {
		response.Error(c, nil, messages.MsgInvalidEmailPassword, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		response.Error(c, nil, messages.MsgInvalidEmailPassword, http.StatusUnauthorized)
		return
	}

	token, err := ctrl.jwtService.GenerateJWT(user.Email)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)

		return
	}

	response.Success(c, gin.H{"token": token}, messages.MsgSuccessful, nil, http.StatusOK)
}
