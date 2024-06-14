package main

import (
	"apcore/controllers"
	"apcore/database"
	"apcore/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.InitDB()
	database.Migrate()
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(middlewares.ErrorHandler())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	auth := router.Group("/auth")
	auth.POST("/signup", controllers.CreateUser)
	auth.POST("/signin", controllers.Login)

	users := router.Group("/users")
	users.GET("/", controllers.GetUsers)

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middlewares.JWTAuthMiddleware())
	adminRoutes.GET("/admin-only", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is an admin-only route"})
	})

	router.Run(":8000")
}
