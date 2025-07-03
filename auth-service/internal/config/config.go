package config

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

type config struct {
	Port      string
	JWTSecret string
	PGConn    string
}

func LoadEnv() config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	get := func(key, def string) string {
		if v := os.Getenv(key); v != "" {
			return v
		}
		return def
	}
	return config{
		Port:      get("PORT", "8000"),
		JWTSecret: get("JWT_SECRET", uuid.NewString()),
		PGConn:    get("PG_CONN", "postgresql://user:pass@postgres:5432/auth"),
	}
}
