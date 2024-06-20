package routes

import (
	"apcore/controllers"
	"apcore/repositories"
	"apcore/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UsersRoutes(router *gin.Engine, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	users := router.Group("/users")
	users.GET("", userController.GetUsers)
}
