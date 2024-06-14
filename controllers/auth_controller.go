package controllers

import (
	"apcore/database"
	"apcore/models"
	"apcore/response"
	"apcore/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
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

	if err := database.GetDB().Create(&user).Error; err != nil {
		response.Error(c, nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Success(c, user, "success", nil, http.StatusCreated)
}

func Login(c *gin.Context) {
	var user models.User
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.GetDB().Where("email = ?", input.Email).First(&user).Error; err != nil {
		response.Error(c, nil, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		response.Error(c, nil, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		response.Error(c, nil, "Failed to generate token", http.StatusInternalServerError)

		return
	}

	response.Success(c, gin.H{"token": token}, "success", nil, http.StatusOK)
}
