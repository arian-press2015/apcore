package middlewares

import (
	"apcore/config"
	"time"

	"github.com/gin-contrib/cors"
)

func GetCorsConfig(cfg *config.Config) cors.Config {
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
