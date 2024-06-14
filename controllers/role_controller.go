package controllers

import (
	"apcore/database"
	"apcore/models"
	"apcore/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	if err := database.GetDB().Create(&role).Error; err != nil {
		response.Error(c, nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Success(c, role, "success", nil, http.StatusOK)
}

func GetRoles(c *gin.Context) {
	var roles []models.Role
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	if err := database.GetDB().Offset(offset).Limit(limit).Preload("Users").Find(&roles).Error; err != nil {
		response.Error(c, nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Success(c, roles, "success", nil, http.StatusOK)
}
