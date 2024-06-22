package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, ctrl *controllers.AuthController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	auth := router.Group("/auth")
	auth.POST("/signup", ctrl.CreateUser)
	auth.POST("/signin", ctrl.Login)
}
