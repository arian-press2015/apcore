package controllers

import (
	"apcore/messages"
	"apcore/models"
	"apcore/response"
	"apcore/services"
	"apcore/utils"
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
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	userID, err := parsers.ParseUUIDFromContext(c, "userID")

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		c.Abort()
		return
	}

	favorites, err := ctrl.service.GetFavorites(offset, limit, userID)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	count, err := ctrl.service.GetFavoritesCount(userID)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	pagination := &response.Pagination{
		Offset: offset,
		Limit:  limit,
		Count:  count,
	}

	response.Success(c, favorites, messages.MsgSuccessful, pagination, http.StatusOK)
}

type CreateFavoriteBody struct {
	CustomerID string `json:"customer_id"`
}

func (ctrl *UserController) AddToFavorites(c *gin.Context) {
	var input CreateFavoriteBody

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, nil, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := parsers.ParseUUIDFromContext(c, "userID")

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		c.Abort()
		return
	}

	customerID, err := utils.UUIDParser(input.CustomerID)
	if err != nil {
		response.Error(c, nil, messages.MsgBadRequest, http.StatusBadRequest)
		c.Abort()
		return
	}

	newFavorite := &models.Favorites{
		UserID:     userID,
		CustomerID: customerID,
	}

	if err := ctrl.service.AddToFavorites(newFavorite); err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, newFavorite, messages.MsgSuccessful, nil, http.StatusCreated)
}

func (ctrl *UserController) DeleteFromFavorites(c *gin.Context) {
	customerIDString := c.Param("customerID")

	userID, err := parsers.ParseUUIDFromContext(c, "userID")

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		c.Abort()
		return
	}

	customerID, err := utils.UUIDParser(customerIDString)
	if err != nil {
		response.Error(c, nil, messages.MsgBadRequest, http.StatusBadRequest)
		c.Abort()
		return
	}

	err = ctrl.service.DeleteFromFavorites(customerID, userID)

	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, nil, messages.MsgSuccessful, nil, http.StatusOK)
}
