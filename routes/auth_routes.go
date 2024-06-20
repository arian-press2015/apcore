package routes

import (
	"apcore/controllers"
	"apcore/repositories"
	"apcore/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	authController := controllers.NewAuthController(userService)

	auth := router.Group("/auth")
	auth.POST("/signup", authController.CreateUser)
	auth.POST("/signin", authController.Login)
}
