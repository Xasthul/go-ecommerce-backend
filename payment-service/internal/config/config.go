package config

import (
	"os"
)

type config struct {
	Port        string
	DatabaseURL string
	RabbitMqUrl string
}

func LoadEnv() config {
	return config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DB_DSN"),
		RabbitMqUrl: os.Getenv("RABBIT_MQ_URL"),
	}
}
