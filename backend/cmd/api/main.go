package main

import (
	"log"

	"app-backend/config"
	"app-backend/internal/delivery/http"
	"app-backend/internal/domain"
	"app-backend/internal/repository/postgres"
	"app-backend/internal/usecase"
	dbpkg "app-backend/pkg/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	// Initialize Database Connection
	db := dbpkg.NewPostgresDB(cfg)

	// Auto-migrate models (optional: recommended using golang-migrate in prod)
	err = db.AutoMigrate(&domain.User{}, &domain.Product{})
	if err != nil {
		log.Fatalf("could not migrate models: %v", err)
	}

	// Setup Repositories
	userRepo := postgres.NewUserRepository(db)
	productRepo := postgres.NewProductRepository(db)

	// Setup Usecases
	userUsecase := usecase.NewUserUsecase(userRepo)
	productUsecase := usecase.NewProductUsecase(productRepo)

	// Setup Echo Router
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize HTTP Handlers
	http.NewHandler(e, userUsecase, productUsecase)

	// Start Server
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
