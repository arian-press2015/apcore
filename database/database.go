package database

import (
	"apcore/config"
	"apcore/models"
	"fmt"
	"log"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Module = fx.Options(
	fx.Provide(NewDB),
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dbConfig := cfg.Database
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbConfig.Host, dbConfig.User, dbConfig.DBName, dbConfig.Password, dbConfig.Port)

	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func Migrate(db *gorm.DB) {
	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error; err != nil {
		log.Fatalf("Failed to create uuid-ossp extension: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Admin{}, &models.Feature{}, &models.Customer{})
}
