package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"geo-app/internal/models"
	"geo-app/internal/services"
)

type WellHandler interface {
	GetWells(c *gin.Context)
	GetWellByID(c *gin.Context)
	CreateWell(c *gin.Context)
	UpdateWell(c *gin.Context)
	DeleteWell(c *gin.Context)
}

type wellHandler struct {
	wellService services.WellService
	validator   *validator.Validate
}

func NewWellHandler(wellService services.WellService) WellHandler {
	validate := validator.New()
	validate.RegisterValidation("latitude", latitudeValidation)
	validate.RegisterValidation("longitude", longitudeValidation)
	validate.RegisterValidation("positive", positiveValidation)
	return &wellHandler{wellService: wellService, validator: validate}
}

func latitudeValidation(fl validator.FieldLevel) bool {
	latitude := fl.Field().Float()
	return latitude >= -90 && latitude <= 90
}

func longitudeValidation(fl validator.FieldLevel) bool {
	longitude := fl.Field().Float()
	return longitude >= -180 && longitude <= 180
}

func positiveValidation(fl validator.FieldLevel) bool {
	value := fl.Field().Float()
	return value >= 0
}

func (h *wellHandler) GetWells(c *gin.Context) {
	wells, err := h.wellService.GetWells()
	if err != nil {
		log.Printf("Error getting wells: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, wells)
}

func (h *wellHandler) GetWellByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Invalid well ID: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	well, err := h.wellService.GetWellByID(uint(id))
	if err != nil {
		log.Printf("Well not found with ID: %d", id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Well not found"})
		return
	}

	c.JSON(http.StatusOK, well)
}

func (h *wellHandler) CreateWell(c *gin.Context) {
	var well models.Well
	if err := c.ShouldBindJSON(&well); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(well); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		log.Printf("Validation error: %v", validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
		return
	}

	createdWell, err := h.wellService.CreateWell(well)
	if err != nil {
		log.Printf("Error creating well: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Well created: %v", createdWell)
	c.JSON(http.StatusCreated, createdWell)
}

func (h *wellHandler) UpdateWell(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Invalid well ID: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var well models.Well
	if err := c.ShouldBindJSON(&well); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(well); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		log.Printf("Validation error: %v", validationErrors)
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErrors.Error()})
		return
	}

	well.ID = uint(id) // Set the ID from the URL parameter

	updatedWell, err := h.wellService.UpdateWell(well)
	if err != nil {
		log.Printf("Error updating well: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Well updated: %v", updatedWell)
	c.JSON(http.StatusOK, updatedWell)
}

func (h *wellHandler) DeleteWell(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("Invalid well ID: %s", idStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	err = h.wellService.DeleteWell(uint(id))
	if err != nil {
		log.Printf("Error deleting well: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Well deleted with ID: %d", id)
	c.Status(http.StatusNoContent)
}
