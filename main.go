package main

import (
	"apcore/config"
	"apcore/database"
	"apcore/logger"
	"apcore/middlewares"
	"apcore/routes"
	"log"

	"github.com/gin-gonic/gin"
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
	defer logger.Sync()

	database.InitDB()
	database.Migrate()
	// acl.InitACL()

	setupMiddlewares(router)

	routes.SetupRoutes(router, database.GetDB())

	err := router.Run(":" + config.AppConfig.Port)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func setupMiddlewares(router *gin.Engine) {
	router.Use(gin.Logger())

	router.Use(middlewares.TrackIdMiddleware())
	router.Use(middlewares.LocaleMiddleware())
	router.Use(middlewares.RecoveryMiddleware())
	router.Use(middlewares.ResponseHandlerMiddleware())
	// router.Use(authz.NewAuthorizer(acl.Enforcer))
}
