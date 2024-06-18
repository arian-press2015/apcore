package database

import (
	"apcore/config"
	"apcore/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB
var err error

func InitDB() {
	dbConfig := config.AppConfig.Database
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbConfig.Host, dbConfig.User, dbConfig.DBName, dbConfig.Password, dbConfig.Port)

	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func GetDB() *gorm.DB {
	return db
}

func Migrate() {
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Admin{})
}
