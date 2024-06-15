package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
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

var AppConfig Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	LoadConfig()
	log.Println("hih")
}

func LoadConfig() {
	AppConfig = Config{}
	// app config
	AppConfig.Port = getEnv("PORT", "8080")
	// database config
	AppConfig.Database.User = getEnv("POSTGRES_USER", "root")
	AppConfig.Database.Password = getEnv("POSTGRES_PASSWORD", "password")
	AppConfig.Database.DBName = getEnv("POSTGRES_DB", "mydatabase")
	AppConfig.Database.Host = getEnv("POSTGRES_HOST", "localhost")
	AppConfig.Database.Port = getEnv("POSTGRES_PORT", "3306")
	// jwt config
	AppConfig.Jwt.JwtSecret = getEnv("JWT_SECRET", "defaultsecret")
	AppConfig.Jwt.JwtExpireAt = getEnv("JWT_EXPIRE_AT", "24h")
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
