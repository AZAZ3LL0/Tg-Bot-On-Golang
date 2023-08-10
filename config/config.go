package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("error loading .env file")
	}
	return os.Getenv(key), nil
}
