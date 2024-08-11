package controllers

import (
	"apcore/messages"
	"apcore/response"
	"apcore/services"
	"apcore/utils/parsers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{service}
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	users, err := ctrl.service.GetUsers(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, users, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *UserController) GetUserById(c *gin.Context) {
	uuidStr := c.Param("uuid")
	uuid, err := parsers.ParseUUID(uuidStr)

	if err != nil {
		response.Error(c, nil, messages.MsgBadRequest, http.StatusBadRequest)
		return
	}

	user, err := ctrl.service.GetUserByID(uuid)
	if err != nil {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
		return
	}

	response.Success(c, user, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *UserController) GetCurrentUser(c *gin.Context) {
	uuid, err := parsers.ParseUUIDFromContext(c, "userID")
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	user, err := ctrl.service.GetUserByID(uuid)
	if err != nil {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
		return
	}

	response.Success(c, user, messages.MsgSuccessful, nil, http.StatusOK)
}

type UpdateProfileBody struct {
	FullName     string `json:"fullName" binding:"required"`
	ProfileImage string `json:"profileImage" binding:"required"`
}

func (ctrl *UserController) UpdateCurrentUser(c *gin.Context) {
	var input UpdateProfileBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	uuid, err := parsers.ParseUUIDFromContext(c, "userID")
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	user, err := ctrl.service.GetUserByID(uuid)
	if err != nil {
		response.Error(c, nil, messages.MsgNotFound, http.StatusNotFound)
		return
	}

	user.FullName = input.FullName
	user.ProfileImage = input.ProfileImage

	err = ctrl.service.UpdateUser(user)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, user, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *UserController) GetFavorites(c *gin.Context) {
	response.Success(c, gin.H{"hi": "bye"}, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *UserController) AddToFavorites(c *gin.Context) {
	response.Success(c, gin.H{"hi": "bye"}, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *UserController) DeleteFromFavorites(c *gin.Context) {
	response.Success(c, gin.H{"hi": "bye"}, messages.MsgSuccessful, nil, http.StatusOK)
}
