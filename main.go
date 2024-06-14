package main

import (
	"apcore/controllers"
	"apcore/database"
	"apcore/middlewares"
	"apcore/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.InitDB()
	database.Migrate()
	router := gin.Default()

	router.Use(gin.Logger())

	router.Use(middlewares.TrackIdMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.ResponseHandlerMiddleware())

	router.Use(middlewares.ErrorHandler())

	routes.SetupRoutes(router)

	router.Run(":8000")
}
