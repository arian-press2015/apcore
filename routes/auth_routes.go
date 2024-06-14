package routes

import (
	"apcore/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	auth.POST("/signup", controllers.CreateUser)
	auth.POST("/signin", controllers.Login)
}