package config

import (
	"log"
	"os"
	"strconv"
)

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

func getEnvAsBool(key string, fallback bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatalf("Error parsing boolean from environment variable %s: %v", key, err)
		}
		return boolValue
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error parsing integer from environment variable %s: %v", key, err)
		}
		return intValue
	}
	return fallback
}

func mustGetEnvAsBool(key string) bool {
	if value, exists := os.LookupEnv(key); exists {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			log.Fatalf("Error parsing boolean from environment variable %s: %v", key, err)
		}
		return boolValue
	}
	log.Fatalf("Environment variable %s not set", key)
	return false
}

func mustGetEnvAsInt(key string) int {
	if value, exists := os.LookupEnv(key); exists {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error parsing integer from environment variable %s: %v", key, err)
		}
		return intValue
	}
	log.Fatalf("Environment variable %s not set", key)
	return 0
}
