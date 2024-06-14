package controllers

import (
	"apcore/database"
	"apcore/models"
	"apcore/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
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

	if err := database.GetDB().Offset(offset).Limit(limit).Preload("Roles").Find(&users).Error; err != nil {
	response.Error(c, nil, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	response.Success(c, users, "success", nil, http.StatusOK)
}
