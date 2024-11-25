package controllers

import (
	"apcore/messages"
	"apcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingMessage struct {
	Message string `json:"message"`
}

type PingController struct{}

func NewPingController() *PingController {
	return &PingController{}
}

// @Summary Healthcheck route
// @Description Responds "ping" with "pong"
// @Tags ping
// @Accept  json
// @Produce  json
// @Param locale header string true "Locale" Enums(en, fa)
// @Success 200 {object} response.SwaggerResponse[PingMessage]
// @Router /ping [get]
func (ctrl *PingController) Ping(c *gin.Context) {
	response.Success(c, gin.H{"message": "pong"}, messages.MsgSuccessful, nil, http.StatusOK)
}
