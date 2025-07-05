package config

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	Port            string
	JWTSecret       string
	DatabaseURL     string
	AccessTokenTTL  int
	RefreshTokenTTL int
}

func LoadEnv() config {
	accessTokenTTL, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TTL"))
	if err != nil {
		log.Fatal("failed to parse ACCESS_TOKEN_TTL: ", err)
	}
	refreshTokenTTL, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TTL"))
	if err != nil {
		log.Fatal("failed to parse REFRESH_TOKEN_TTL")
	}

	return config{
		Port:            os.Getenv("PORT"),
		JWTSecret:       os.Getenv("JWT_SECRET"),
		DatabaseURL:     os.Getenv("DB_DSN"),
		AccessTokenTTL:  accessTokenTTL,
		RefreshTokenTTL: refreshTokenTTL,
	}
}
