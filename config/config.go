package config

import (
	"log"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(NewConfig),
)

type Config struct {
	Env      string
	Port     string
	Database struct {
		User     string
		Password string
		DBName   string
		Host     string
		Port     string
	}
	Logger struct {
		ElasticEnabled bool
		FileEnabled    bool
	}
	Elastic struct {
		Url      string
		Username string
		Password string
		Index    string
	}
	Redis struct {
		Url string
	}
	Jwt struct {
		JwtSecret   string
		JwtExpireAt string
	}
	Sms struct {
		ApiUrl     string
		ApiKey     string
		LineNumber string
	}
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	config := &Config{}
	// app config
	config.Env = getEnv("ENV", "development")
	config.Port = getEnv("PORT", "8080")
	// database config
	config.Database.User = getEnv("POSTGRES_USER", "root")
	config.Database.Password = getEnv("POSTGRES_PASSWORD", "password")
	config.Database.DBName = getEnv("POSTGRES_DB", "mydatabase")
	config.Database.Host = getEnv("POSTGRES_HOST", "localhost")
	config.Database.Port = getEnv("POSTGRES_PORT", "5432")
	// logger config
	config.Logger.ElasticEnabled = getEnvAsBool("LOGGER_ELASTIC_ENABLED", false)
	config.Logger.FileEnabled = getEnvAsBool("LOGGER_FILE_ENABLED", true)
	// elasticsearch config
	config.Elastic.Url = getEnv("ELASTICSEARCH_URL", "https://localhost:9200")
	config.Elastic.Username = getEnv("ELASTICSEARCH_USERNAME", "elastic")
	config.Elastic.Password = getEnv("ELASTICSEARCH_PASSWORD", "password")
	config.Elastic.Index = getEnv("ELASTICSEARCH_INDEX", "apcore_logs")
	// redis config
	config.Redis.Url = getEnv("REDIS_URL", "127.0.0.1:6379")
	// jwt config
	config.Jwt.JwtSecret = getEnv("JWT_SECRET", "defaultsecret")
	config.Jwt.JwtExpireAt = getEnv("JWT_EXPIRE_AT", "24h")
	// sms config
	config.Sms.ApiUrl = mustGetEnv("SMS_API_URL")
	config.Sms.ApiKey = mustGetEnv("SMS_API_KEY")
	config.Sms.LineNumber = mustGetEnv("SMS_LINE_NUMBER")

	return config
}
