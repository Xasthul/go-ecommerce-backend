package main

import (
	"context"
	"log"

	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/config"
	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/handler"
	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/repository"
	gen "github.com/Xasthul/go-ecommerce-backend/payment-service/internal/repository/db/gen"
	"github.com/Xasthul/go-ecommerce-backend/payment-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := config.LoadEnv()

	// rabbitConn, _ := amqp.Dial(cfg.RabbitMqUrl)
	// pubslisher, err := rabbitmq.NewPublisher(rabbitConn)
	// if err != nil {
	// 	log.Fatal("connect rabbitmq: ", err)
	// }
	// defer rabbitConn.Close()

	databaseURL := cfg.DatabaseURL
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("connect postgres: ", err)
	}
	defer db.Close()
	runMigrations(databaseURL)

	queries := gen.New(db)
	paymentRepository := repository.NewPaymentRepository(queries)
	paymentService := service.NewOrderService(paymentRepository)
	apiHandler := handler.NewApiHandler(paymentService)

	r := gin.Default()
	r.Use(gin.Recovery())
	apiHandler.RegisterRoutes(r)
	r.Run(":" + cfg.Port)
}

func runMigrations(databaseURL string) {
	m, err := migrate.New(
		"file://internal/repository/db/migrations",
		databaseURL,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("failed to run migrations: %v", err)
	}
	log.Println("database migrated successfully")
}
