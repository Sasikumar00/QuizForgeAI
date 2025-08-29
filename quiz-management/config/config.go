package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	MongoURL string
	DBName   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env found")
	}

	return &Config {
		Port: getEnv("PORT", "5000"),
		MongoURL: getEnv("MONGOURL", "mongodb://localhost:27017"),
		DBName: getEnv("DBNAME", "quizforge"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}