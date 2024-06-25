package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, ctrl *controllers.AuthController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	auth := router.Group("/auth")
	auth.POST("/", ctrl.Auth)
	auth.POST("/verify", ctrl.VerifyAuth)
}
