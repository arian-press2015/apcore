package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func RolesRoutes(router *gin.Engine) {
	roles := router.Group("/")
	roles.Use(middlewares.JWTAuthMiddleware())
	roles.POST("/roles", controllers.CreateRole)
	roles.GET("/roles", controllers.GetRoles)
}