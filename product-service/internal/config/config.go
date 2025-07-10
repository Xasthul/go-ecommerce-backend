package config

import (
	"os"
)

type config struct {
	Port          string
	DatabaseURL   string
	ApiKey        string
	RabbitMqUrl   string
	RedisAddress  string
	RedisPassword string
}

func LoadEnv() config {
	return config{
		Port:          os.Getenv("PORT"),
		DatabaseURL:   os.Getenv("DB_DSN"),
		ApiKey:        os.Getenv("API_KEY"),
		RabbitMqUrl:   os.Getenv("RABBIT_MQ_URL"),
		RedisAddress:  os.Getenv("REDIS_ADDR"),
		RedisPassword: os.Getenv("REDIS_PASS"),
	}
}
