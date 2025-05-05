package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBPort       string
	MaxOpenConns int
	MaxIdleConns int
	FrontendUrl  string
}

var Env Config

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	Env = Config{
		Port:         getEnv("APP_PORT", "8080"),
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", ""),
		DBName:       getEnv("DB_NAME", "postgres"),
		DBPort:       getEnv("DB_PORT", "5432"),
		MaxOpenConns: getEnvAsInt("MAX_OPEN_CONNECTINOS", 10),
		MaxIdleConns: getEnvAsInt("MAX_IDLE_CONNECTIONS", 5),
		FrontendUrl: getEnv("FRONTEND_URL", "http://localhost:3000"),
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if valueStr, exists := os.LookupEnv(key); exists {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
		log.Printf("Invalid value for %s. Using fallback: %d\n", key, fallback)
	}
	return fallback
}

func getEnvAsInt32(key string, fallback int32) int32 {
	if valueStr, exists := os.LookupEnv(key); exists {
		value, err := strconv.ParseInt(valueStr, 10, 32);
		if err == nil {
			return int32(value)
		}
		log.Printf("Invalid value for %s. Using fallback: %d\n", key, fallback)
	}
	return fallback
}