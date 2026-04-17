package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
    PostgreDBHost       string
    PostgreDBPort     string
    PostgreDBUser     string
    PostgreDBPassword string
}

func Load() *Config {
    err := godotenv.Load()
    if err != nil {
        // Don't fatal here — in production, env vars come from the system
        log.Println("No .env file found, reading from environment")
    }

    return &Config{
        PostgreDBHost:     getEnv("PG_DB_HOST", "8080"),         // second arg = default
        PostgreDBPort:     getEnv("PG_DB_PORT", "localhost"),
        PostgreDBUser:     getEnv("PG_DB_USER", ""),
        PostgreDBPassword: getEnv("PG_DB_PASSWORD", ""),
    }
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}