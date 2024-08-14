package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func UsersRoutes(router *gin.Engine, ctrl *controllers.UserController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	users := router.Group("/users")
	users.Use(jwtAuthMiddleware.Middleware())
	users.GET("", ctrl.GetUsers)

	// profile
	users.GET("profile", ctrl.GetCurrentUser)
	users.PUT("profile", ctrl.UpdateCurrentUser)

	// favorites
	users.GET("favorites", ctrl.GetFavorites)
	users.POST("favorites", ctrl.AddToFavorites)
	users.DELETE("favorites/:customerID", ctrl.DeleteFromFavorites)

	users.GET(":uuid", ctrl.GetUserById)
}
