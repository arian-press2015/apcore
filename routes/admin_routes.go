package routes

import (
	"apcore/middlewares"
	"apcore/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middlewares.JWTAuthMiddleware())
	adminRoutes.GET("/admin-only", func(c *gin.Context) {
		response.Success(c, nil, "This is an admin-only route", nil, http.StatusOK)
	})
}
