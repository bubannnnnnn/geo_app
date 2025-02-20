package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"geo-app/internal/config"
	"geo-app/internal/handlers"
	"geo-app/internal/models"
	"geo-app/internal/repositories"
	"geo-app/internal/services"
)

func main() {
	// Load configuration from .env file
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Database configuration
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// AutoMigrate database schema
	db.AutoMigrate(&models.Well{}, &models.User{}, &models.Sample{}) // Migrate Sample model

	// Initialize repositories, services, and handlers
	wellRepository := repositories.NewWellRepository(db)
	wellService := services.NewWellService(wellRepository)
	wellHandler := handlers.NewWellHandler(wellService)

	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	sampleRepository := repositories.NewSampleRepository(db)
	sampleService := services.NewSampleService(sampleRepository)
	sampleHandler := handlers.NewSampleHandler(sampleService)

	// Gin setup
	router := gin.Default()

	// Custom validation function
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("latitude", func(fl validator.FieldLevel) bool {
			latitude := fl.Field().Float()
			return latitude >= -90 && latitude <= 90
		})
		v.RegisterValidation("longitude", func(fl validator.FieldLevel) bool {
			longitude := fl.Field().Float()
			return longitude >= -180 && longitude <= 180
		})
	}

	// Serve static files from the "web" directory
	router.Static("/web", "./web")

	// Serve the index.html file at the root
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/web/index.html")
	})

	// Authentication routes
	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	// Well routes
	router.GET("/api/wells", wellHandler.GetWells)
	router.POST("/api/wells", wellHandler.CreateWell)
	router.GET("/api/wells/:id", wellHandler.GetWellByID)
	router.PUT("/api/wells/:id", wellHandler.UpdateWell)
	router.DELETE("/api/wells/:id", wellHandler.DeleteWell)

	// Sample routes
	router.POST("/api/samples", sampleHandler.CreateSample)
	router.GET("/api/samples/:id", sampleHandler.GetSampleByID)
	router.PUT("/api/samples/:id", sampleHandler.UpdateSample)
	router.DELETE("/api/samples/:id", sampleHandler.DeleteSample)
	router.GET("/api/samples/well/:well_id", sampleHandler.GetSamplesByWellID)

	// Start the server
	port := cfg.ServerPort
	log.Printf("Server listening on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
