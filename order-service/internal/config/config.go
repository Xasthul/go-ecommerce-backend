package config

import (
	"os"
)

type config struct {
	Port                 string
	DatabaseURL          string
	ProductServiceUrl    string
	ProductServiceApiKey string
}

func LoadEnv() config {
	return config{
		Port:                 os.Getenv("PORT"),
		DatabaseURL:          os.Getenv("DB_DSN"),
		ProductServiceUrl:    os.Getenv("PRODUCT_SERVICE_URL"),
		ProductServiceApiKey: os.Getenv("PRODUCT_SERVICE_API_KEY"),
	}
}
