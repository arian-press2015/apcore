package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func RolesRoutes(router *gin.Engine, ctrl *controllers.RoleController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	roles := router.Group("/")
	roles.Use(jwtAuthMiddleware.Middleware())
	roles.POST("/roles", ctrl.CreateRole)
	roles.GET("/roles", ctrl.GetRoles)
}
