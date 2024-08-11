package controllers

import (
	"apcore/messages"
	"apcore/response"
	"apcore/services"
	"apcore/utils/parsers"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificationController struct {
	service services.NotificationService
}

func NewNotificationController(service services.NotificationService) *NotificationController {
	return &NotificationController{service}
}

func (ctrl *NotificationController) GetNotifications(c *gin.Context) {
	offset, limit := parsers.ParsePaginationParams(c.Query("offset"), c.Query("limit"))

	notifications, err := ctrl.service.GetNotifications(offset, limit)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, notifications, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *NotificationController) MarkAsRead(c *gin.Context) {
	uuidStr := c.Param("uuid")
	uuid, err := parsers.ParseUUID(uuidStr)

	if err != nil {
		response.Error(c, nil, messages.MsgBadRequest, http.StatusBadRequest)
		return
	}

	err = ctrl.service.MarkAsRead(uuid)
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, nil, messages.MsgSuccessful, nil, http.StatusOK)
}

func (ctrl *NotificationController) MarkAllAsRead(c *gin.Context) {
	err := ctrl.service.MarkAllAsRead()
	if err != nil {
		response.Error(c, nil, messages.MsgInternalServerError, http.StatusInternalServerError)
		return
	}

	response.Success(c, nil, messages.MsgSuccessful, nil, http.StatusOK)
}
