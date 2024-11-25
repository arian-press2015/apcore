package database

import (
	"apcore/config"
	"fmt"
	"time"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	MAX_IDLE_CONNS    = 10
	MAX_OPEN_CONNS    = 100
	CONN_MAX_LIFETIME = 1 * time.Hour
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

	sqlDB.SetMaxIdleConns(MAX_IDLE_CONNS)
	sqlDB.SetMaxOpenConns(MAX_OPEN_CONNS)
	sqlDB.SetConnMaxLifetime(CONN_MAX_LIFETIME)

	return db, nil
}
