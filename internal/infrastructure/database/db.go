package database

import (
	"backend/internal/config"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		config.Env.DBHost,
		config.Env.DBUser,
		config.Env.DBPassword,
		config.Env.DBName,
		config.Env.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from gorm: %v", err)
	}
	sqlDB.SetMaxOpenConns(config.Env.MaxOpenConns)
	sqlDB.SetMaxIdleConns(config.Env.MaxIdleConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}