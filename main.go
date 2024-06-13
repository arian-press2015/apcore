package main

import (
	"apcore/controllers"
	"apcore/database"
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

	router.POST("/users", controllers.CreateUser)
	router.POST("/login", controllers.Login)

	router.Run(":8000")
}
