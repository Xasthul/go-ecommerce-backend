package config

import (
	"os"
)

type config struct {
	Port        string
	JWTSecret   string
	DatabaseURL string
}

func LoadEnv() config {
	return config{
		Port:        os.Getenv("PORT"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
		DatabaseURL: os.Getenv("DB_DSN"),
	}
}
