package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/Xasthul/go-ecommerce-backend/auth-service/internal/config"
	"github.com/Xasthul/go-ecommerce-backend/auth-service/internal/handler"
	"github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository"
	gen "github.com/Xasthul/go-ecommerce-backend/auth-service/internal/repository/db/gen"
	"github.com/Xasthul/go-ecommerce-backend/auth-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/mattes/migrate"
)

func main() {
	cfg := config.LoadEnv()

	databaseURL := cfg.PGConn

	db, err := sql.Open("pgx", cfg.PGConn)
	if err != nil {
		log.Fatal("connect postgres")
	}
	db.SetConnMaxIdleTime(time.Minute)

	queries := gen.New(db)
	userRepository := repository.NewUserRepository(queries)
	authService := service.NewAuthService(userRepository)
	apiHandler := handler.NewAPIHandler(authService)

	runMigrations(databaseURL)

	r := gin.Default()

	r.POST("/register", apiHandler.RegisterHandler)
	r.POST("/login", apiHandler.LoginHandler)

	r.Run(":" + cfg.Port)
}

func runMigrations(databaseURL string) {
	m, err := migrate.New(
		"file://db/migration",
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
