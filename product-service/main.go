package main

import (
	"context"
	"log"

	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/config"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/handler"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/rabbitmq"
	redisdb "github.com/Xasthul/go-ecommerce-backend/product-service/internal/redis"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository"
	gen "github.com/Xasthul/go-ecommerce-backend/product-service/internal/repository/db/gen"
	"github.com/Xasthul/go-ecommerce-backend/product-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	cfg := config.LoadEnv()

	redisClient, err := redisdb.InitRedis(cfg.RedisAddress, cfg.RedisPassword)
	if err != nil {
		log.Fatal("connect redis: ", err)
	}

	databaseURL := cfg.DatabaseURL
	db, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatal("connect postgres: ", err)
	}
	defer db.Close()
	runMigrations(databaseURL)

	queries := gen.New(db)
	productRepository := repository.NewProductRepository(queries)
	categoryRepository := repository.NewCategoryRepository(queries)
	productService := service.NewProductService(productRepository, redisClient)
	categorytService := service.NewCategoryService(categoryRepository)
	apiHandler := handler.NewApiHandler(productService, categorytService)

	rabbitConn, err := amqp.Dial(cfg.RabbitMqUrl)
	if err != nil {
		log.Fatal("connect rabbitmq: ", err)
	}
	defer rabbitConn.Close()
	consumeRabbitMqEvents(rabbitConn, productService)

	r := gin.Default()
	r.Use(gin.Recovery())
	apiHandler.RegisterRoutes(r, cfg.ApiKey)
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

func consumeRabbitMqEvents(rabbitConn *amqp.Connection, productService *service.ProductService) {
	err := rabbitmq.ConsumeOrders(
		rabbitConn,
		func(event *rabbitmq.OrderCreatedEvent) {
			productService.DecreaseStock(
				context.Background(),
				event.ProductID,
				event.Quantity,
			)
		})
	if err != nil {
		log.Fatal("consumer orders: ", err)
	}
}
