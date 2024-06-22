package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminAuthRoutes(router *gin.Engine, ctrl *controllers.AdminAuthController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	auth := router.Group("/admin")
	auth.POST("/signin", ctrl.AdminLogin)
}
