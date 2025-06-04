package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	HTTPPort   string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		HTTPPort:   os.Getenv("HTTP_PORT"),
	}
}
func GetDatabaseURL(config Config) string {
	return fmt.Sprintf("host=%s user=%s password=%s port=%s connect_timeout=60 dbname=%s sslmode=disable search_path=salary",
		config.DBHost, config.DBUser, config.DBPassword, config.DBPort, config.DBName)
}
