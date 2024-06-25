package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminAuthRoutes(router *gin.Engine, ctrl *controllers.AdminAuthController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	auth := router.Group("/admin")
	auth.POST("/auth", ctrl.AdminLogin)
	auth.POST("/auth/enable2fa", ctrl.Enable2FA)
	auth.POST("/auth/verify2fa", ctrl.Verify2FA)
}
