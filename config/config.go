package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	Domain string
}

var Env Config

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, relying on environment variables")
	}

	Env = Config{
		Port:   getEnv("PORT", "8080"),
		Domain: getEnv("DOMAIN", "localhost"),
	}
	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}