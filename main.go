package main

import (
	"apcore/config"
	"apcore/database"
	"apcore/middlewares"
	"apcore/routes"
	"log"

	_ "apcore/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title APCore API
// @version 0.1
// @description This is the core of AP2015 projects
// @termsOfService http://your_project/terms/

// @contact.name AP2015
// @contact.url http://www.your_project.com/support
// @contact.email arian.press2015@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err := router.Run(":" + config.AppConfig.Port)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
