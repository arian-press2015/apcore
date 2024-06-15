package main

import (
	"apcore/config"
	"apcore/database"
	"apcore/middlewares"
	"apcore/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.InitDB()
	database.Migrate()
	router := gin.Default()

	router.Use(gin.Logger())

	router.Use(middlewares.TrackIdMiddleware())
	router.Use(middlewares.LocaleMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.ResponseHandlerMiddleware())

	router.Use(middlewares.ErrorHandler())

	routes.SetupRoutes(router)
	err := router.Run(":" + config.AppConfig.Port)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
