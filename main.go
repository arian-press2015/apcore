package main

import (
	"apcore/config"
	"apcore/controllers"
	"apcore/database"
	"apcore/logger"
	"apcore/middlewares"
	"apcore/repositories"
	"apcore/routes"
	"apcore/services"
	"apcore/utils"
	"context"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
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
	app := fx.New(
		config.Module,
		database.Module,
		logger.Module,
		utils.Module,
		routes.Module,
		middlewares.Module,
		repositories.Module,
		services.Module,
		controllers.Module,
		fx.Invoke(registerHooks),
	)

	app.Run()
}

func registerHooks(
	lc fx.Lifecycle,
	router *gin.Engine,
	middlewares *middlewares.Middlewares,
	routes *routes.Routes,
	cfg *config.Config,
	logger *logger.Logger,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			middlewares.SetupMiddlewares(router)
			routes.SetupRoutes(router)
			go func() {
				if err := router.Run(":" + cfg.Port); err != nil {
					logger.Fatal("Server failed to start", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Sync()
			return nil
		},
	})
}
