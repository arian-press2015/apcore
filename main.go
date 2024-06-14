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

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", controllers.CreateUser)
	router.POST("/login", controllers.Login)

	protected := router.Group("/")
	protected.Use(middlewares.JWTAuthMiddleware())
	protected.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is a protected route"})
	})

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middlewares.JWTAuthMiddleware())
	adminRoutes.GET("/admin-only", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "This is an admin-only route"})
	})

	router.Run(":8000")
}
