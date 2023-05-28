package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv - функция для загрузки переменных окружения из файла .env
func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	return nil
}

// GetEnv - функция для получения значения переменной окружения
func GetEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Env variable %s not set", key)
	}

	return value
}
