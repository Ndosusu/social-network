package configs

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "localhost:5432"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}