package database

import (
	"apcore/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB
var err error

func InitDB() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", dbHost, dbUser, dbName, dbPassword, dbPort)

	db, err = gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to database %v", err)
	}

	db.LogMode(true)
}

func GetDB() *gorm.DB {
	return db
}

func Migrate() {
	db.AutoMigrate(&models.User{})
}
