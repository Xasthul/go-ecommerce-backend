package config

import (
	"os"
)

type config struct {
	Port              string
	AuthServiceURL    string
	ProductServiceURL string
	OrderServiceURL   string
	PaymentServiceURL string
	JwtSecret         string
}

func LoadEnv() config {
	return config{
		Port:              os.Getenv("PORT"),
		AuthServiceURL:    os.Getenv("AUTH_SERVICE_URL"),
		ProductServiceURL: os.Getenv("PRODUCT_SERVICE_URL"),
		OrderServiceURL:   os.Getenv("ORDER_SERVICE_URL"),
		JwtSecret:         os.Getenv("JWT_SECRET"),
		PaymentServiceURL: os.Getenv("PAYMENT_SERVICE_URL"),
	}
}
