package config

import (
	"os"
)

type config struct {
	Port           string
	AuthServiceURL string
}

func LoadEnv() config {
	return config{
		Port:           os.Getenv("PORT"),
		AuthServiceURL: os.Getenv("AUTH_SERVICE_URL"),
	}
}
