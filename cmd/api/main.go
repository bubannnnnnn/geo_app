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
	// Загрузка данных из .env
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Конфигурация бд
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

	// Миграция бд
	db.AutoMigrate(&models.Well{}, &models.User{}, &models.Sample{}, &models.Material{})

	// Утановка repositories, services, и handlers
	wellRepository := repositories.NewWellRepository(db)
	wellService := services.NewWellService(wellRepository)
	wellHandler := handlers.NewWellHandler(wellService)

	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	sampleRepository := repositories.NewSampleRepository(db)
	sampleService := services.NewSampleService(sampleRepository)
	sampleHandler := handlers.NewSampleHandler(sampleService)

	materialRepository := repositories.NewMaterialRepository(db)
	materialService := services.NewMaterialService(materialRepository)
	materialHandler := handlers.NewMaterialHandler(materialService)

	// Gin установка
	router := gin.Default()

	// Функция кастомного фалидатора
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// Валидатор
		v.RegisterValidation("latitude", func(fl validator.FieldLevel) bool {
			latitude := fl.Field().Float()
			return latitude >= -90 && latitude <= 90
		})

		// Longitude validation
		v.RegisterValidation("longitude", func(fl validator.FieldLevel) bool {
			longitude := fl.Field().Float()
			return longitude >= -180 && longitude <= 180
		})

		// Positive number validation
		v.RegisterValidation("positive", func(fl validator.FieldLevel) bool {
			value := fl.Field().Float()
			return value >= 0
		})
	}

	// Обслуживать статические файлы из каталога "web
	router.Static("/web", "./web")

	// Обработать файл index.html в корневом каталоге
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/web/index.html")
	})

	// Добавление маршрута для material_page.html
	router.GET("/material", func(c *gin.Context) {
		c.File("./web/material.html")
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

	// Material routes
	router.GET("/api/materials", materialHandler.GetAllMaterials)
	router.GET("/api/materials/:id", materialHandler.GetMaterialByID)
	router.POST("/api/materials", materialHandler.CreateMaterial)
	router.PUT("/api/materials/:id", materialHandler.UpdateMaterial)
	router.DELETE("/api/materials/:id", materialHandler.DeleteMaterial)

	// Старт сервер
	port := cfg.ServerPort
	log.Printf("Server listening on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}
