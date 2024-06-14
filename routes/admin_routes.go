package routes

import (
	"apcore/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middlewares.JWTAuthMiddleware())
	adminRoutes.GET("/admin-only", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is an admin-only route"})
	})
}
