package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine, ctrl *controllers.UserController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	users := router.Group("/users")
	users.GET("", ctrl.GetUsers)
	users.GET(":uuid", ctrl.GetUserById)
}
