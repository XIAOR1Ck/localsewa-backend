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
		JWTSecret 					string
		JWTIssuer 					string
		JWTTTL 							time.Duration
}

func Load() *Config {
    // Load .env file if it exists
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, using environment variables")
    }

    return &Config{
        ServerPort:          getEnv("SERVER_PORT", "8081"),
      	DatabaseURL:         getEnv("GOOSE_DBSTRING", ""),
        FirebaseCredentials: getEnv("FIREBASE_CREDENTIALS", ""),
        Environment:         getEnv("ENVIRONMENT", "development"),
        LoggerLevel:         getEnv("LOGGER_LEVEL", "info"),
				JWTSecret: 					 getEnv("JWT_SECRET", "super_secret_jwt_string"),
				JWTIssuer: 					 getEnv("JWT_ISSUER", "LocalSewa"),
				JWTTTL: 						 getEnvAsDuration("JWT_TTL", 7*24*time.Hour),
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
