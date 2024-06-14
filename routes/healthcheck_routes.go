package routes

import (
	"apcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		response.Success(c, nil, "pong", nil, http.StatusOK)
	})
}
