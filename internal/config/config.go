package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jonathanmoreiraa/planejja/pkg/log"
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
		log.NewLogger().Error(err)
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
