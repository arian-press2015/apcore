package routes

import (
	"apcore/controllers"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	users.GET("", controllers.GetUsers)
}