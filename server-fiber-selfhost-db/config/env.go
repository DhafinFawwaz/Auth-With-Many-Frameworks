package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetEnv(key string) string {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			fmt.Printf("Error loading .env file: %v\n", err)
		}
	}

	value := os.Getenv(key)
	if value == "" {
		fmt.Printf("Environment variable %s is not set\n", key)
	}
	return value
}
