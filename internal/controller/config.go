package controller

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	TokenSecret     []byte
	TokenExpiration int
	DBPath          string
}

func DefaultConfig() *Config {
	godotenv.Load(".env")
	tokenExpiration, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HRS"))

	return &Config{
		Port:            os.Getenv("PORT"),
		TokenSecret:     []byte(os.Getenv("TOKEN_SECRET")),
		TokenExpiration: tokenExpiration,
		DBPath:          os.Getenv("DB_PATH"),
	}
}
