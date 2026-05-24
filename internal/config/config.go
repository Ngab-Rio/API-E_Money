package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppPort string
	GinMode string

	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	DBSSLMode string

	JWTSecret string
	JWTIssuer string
}

func Load() *Config {
	if err := godotenv.Load(); err != nil {
		log.Println(".env Tidak Ada Dalam Folder Proyek")
	}

	cfg := &Config{
		AppPort: getEnv("APP_PORT", "8080"),
		GinMode: getEnv("GIN_MODE", "debug"),

		DBHost:    getEnv("DB_HOST", "localhost"),
		DBPort:    getEnv("DB_PORT", "5432"),
		DBUser:    getEnv("DB_USER", "postgres"),
		DBPass:    getEnv("DB_PASSWORD", ""),
		DBName:    getEnv("DB_NAME", ""),
		DBSSLMode: getEnv("DB_SSLMODE", "disable"),
	}

	return cfg
}

func (cg *Config) DBDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cg.DBHost,
		cg.DBPort,
		cg.DBUser,
		cg.DBPass,
		cg.DBName,
		cg.DBSSLMode,
	)
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
