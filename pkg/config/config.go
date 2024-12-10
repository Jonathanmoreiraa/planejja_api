package config

import (
	"fmt"
	"os"
)

// Config estrutura para armazenar as configurações da aplicação.
type Config struct {
	AppName    string
	DbUser     string
	DbPassword string
	DbName     string
	DbHost     string
	DbPort     string
}

// LoadConfig carrega as configurações do arquivo .env ou variáveis de ambiente.
func LoadConfig() (*Config, error) {
	if appNameEnv := os.Getenv("APP_NAME"); appNameEnv == "" {
		return nil, fmt.Errorf("Arquivo .env não encontrado, usando variáveis de ambiente")
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

func (c Config) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		c.DbUser, c.DbPassword, c.DbHost, c.DbPort, c.DbName,
	)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
