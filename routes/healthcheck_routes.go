package routes

import (
	"apcore/messages"
	"apcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		response.Success(c, gin.H{"message": "pong"}, messages.MsgSuccessful, nil, http.StatusOK)
	})
}
