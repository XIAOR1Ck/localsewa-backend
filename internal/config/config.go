package config

import (
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
)

type Config struct {
    ServerPort          string
    DatabaseURL         string
    FirebaseCredentials string
    Environment         string
    LoggerLevel         string
}

func Load() *Config {
    // Load .env file if it exists
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    return &Config{
        ServerPort:          getEnv("SERVER_PORT", "8081"),
        DatabaseURL:         getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/yourdb?sslmode=disable"),
        FirebaseCredentials: getEnv("FIREBASE_CREDENTIALS", ""),
        Environment:         getEnv("ENVIRONMENT", "development"),
        LoggerLevel:         getEnv("LOGGER_LEVEL", "info"),
    }
}

func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}

// Helper to parse duration from env
func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
    if value := os.Getenv(key); value != "" {
        if duration, err := time.ParseDuration(value); err == nil {
            return duration
        }
    }
    return defaultValue
}
