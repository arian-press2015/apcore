package middlewares

import (
	"apcore/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware(cfg *config.Config) gin.HandlerFunc {
	corsConfig := getCorsConfig(cfg)
	return cors.New(corsConfig)
}

func getCorsConfig(cfg *config.Config) cors.Config {
	corsConfig := cors.Config{
		AllowOrigins:     []string{cfg.CorsEndpoint},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	return corsConfig
}
