package server

import (
	"os"
	"strconv"

	"github.com/arsu4ka/todo-auth/internal/dbs"
	"github.com/joho/godotenv"
)

type Config struct {
	Port            string
	BaseURL         string
	SMTPEmail       string
	SMTPPassword    string
	TokenSecret     string
	TokenExpiration int
	DBConf          *dbs.Config
}

func DefaultConfig() *Config {
	godotenv.Load(".env")
	tokenExpiration, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HRS"))

	return &Config{
		Port:            os.Getenv("PORT"),
		BaseURL:         os.Getenv("URL"),
		SMTPEmail:       os.Getenv("SMTP_EMAIL"),
		SMTPPassword:    os.Getenv("SMTP_PASSWORD"),
		TokenSecret:     os.Getenv("TOKEN_SECRET"),
		TokenExpiration: tokenExpiration,
		DBConf: &dbs.Config{
			Host:     os.Getenv("PGHOST"),
			Port:     os.Getenv("PGPORT"),
			Name:     os.Getenv("PGDATABASE"),
			User:     os.Getenv("PGUSER"),
			Password: os.Getenv("PGPASSWORD"),
		},
	}
}
