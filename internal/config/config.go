package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppName    string
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if appNameEnv := os.Getenv("APP_NAME"); appNameEnv == "" {
		return nil, fmt.Errorf("arquivo .env não encontrado, usando variáveis de ambiente")
	}

	return &Config{
		AppName:    getEnv("APP_NAME", ""),
		DbPort:     getEnv("DB_PORT", "8080"),
		DbUser:     getEnv("DB_USER", "root"),
		DbPassword: getEnv("DB_PASSWORD", "password"),
		DbName:     getEnv("DB_NAME", "mydatabase"),
		DbHost:     getEnv("DB_HOST", "localhost"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
