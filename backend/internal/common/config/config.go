package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Environment              string
	Port                     string
	PostgresConnectionString string
	MongoConnectionString    string
	MongoUsername            string
	MongoPassword            string
	MongoDatabase            string
}

var cfg Config

func Init(path string) {
	err := godotenv.Load(path)
	if err != nil {
		log.Fatal(err)
	}

	cfg.Environment = getEnv("ENVIRONMENT", "development")
	cfg.Port = getEnv("PORT", "8080")
	cfg.PostgresConnectionString = getEnv("DB_CONNECTION_STRING", "")
	cfg.MongoConnectionString = getEnv("NOSQL_CONNECTION_STRING", "")
	cfg.MongoUsername = getEnv("NOSQL_USERNAME", "mongo")
	cfg.MongoPassword = getEnv("NOSQL_PASSWORD", "mongo")
	cfg.MongoDatabase = getEnv("NOSQL_DATABASE", "family-planner")
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	return defaultValue
}

func GetConfig() Config {
	return cfg
}
