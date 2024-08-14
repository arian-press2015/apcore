package routes

import (
	"apcore/controllers"
	"apcore/middlewares"

	"github.com/gin-gonic/gin"
)

func CommentsRoutes(router *gin.Engine, ctrl *controllers.CommentController, jwtAuthMiddleware *middlewares.JWTAuthMiddleware) {
	comments := router.Group("/")
	comments.Use(jwtAuthMiddleware.Middleware())
	comments.POST("/comments", ctrl.CreateComment)
	comments.GET("/comments", ctrl.GetComments)
}
