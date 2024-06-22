package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
)

type Config struct {
	Port     string
	Database struct {
		User     string
		Password string
		DBName   string
		Host     string
		Port     string
	}
	Jwt struct {
		JwtSecret   string
		JwtExpireAt string
	}
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	config := &Config{}
	// app config
	config.Port = getEnv("PORT", "8080")
	// database config
	config.Database.User = getEnv("POSTGRES_USER", "root")
	config.Database.Password = getEnv("POSTGRES_PASSWORD", "password")
	config.Database.DBName = getEnv("POSTGRES_DB", "mydatabase")
	config.Database.Host = getEnv("POSTGRES_HOST", "localhost")
	config.Database.Port = getEnv("POSTGRES_PORT", "5432")
	// jwt config
	config.Jwt.JwtSecret = getEnv("JWT_SECRET", "defaultsecret")
	config.Jwt.JwtExpireAt = getEnv("JWT_EXPIRE_AT", "24h")

	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func mustGetEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Fatalf("Environment variable %s not set", key)
	return ""
}
