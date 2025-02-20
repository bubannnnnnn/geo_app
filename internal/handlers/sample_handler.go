package handlers

import (
	"fmt"
	"geo-app/internal/models"
	"geo-app/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SampleHandler struct {
	sampleService *services.SampleService
}

func NewSampleHandler(sampleService *services.SampleService) *SampleHandler {
	return &SampleHandler{sampleService: sampleService}
}

// Create Sample
func (h *SampleHandler) CreateSample(c *gin.Context) {
	var sample models.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.sampleService.CreateSample(&sample); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create sample"})
		return
	}

	c.JSON(http.StatusCreated, sample)
}

// Get Sample by ID
func (h *SampleHandler) GetSampleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sample ID"})
		return
	}

	sample, err := h.sampleService.GetSampleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sample not found"})
		return
	}

	c.JSON(http.StatusOK, sample)
}

// Update Sample
func (h *SampleHandler) UpdateSample(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sample ID"})
		return
	}

	var sample models.Sample
	if err := c.ShouldBindJSON(&sample); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sample.ID = uint(id) // Ensure the ID is set for update
	if err := h.sampleService.UpdateSample(&sample); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update sample"})
		return
	}

	c.JSON(http.StatusOK, sample)
}

// Delete Sample
func (h *SampleHandler) DeleteSample(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sample ID"})
		return
	}

	if err := h.sampleService.DeleteSample(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete sample"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Sample deleted"})
}

// Get Samples by WellID (Added)
func (h *SampleHandler) GetSamplesByWellID(c *gin.Context) {
	wellIDStr := c.Param("well_id")
	wellID, err := strconv.Atoi(wellIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid well ID"})
		return
	}

	samples, err := h.sampleService.GetSamplesByWellID(uint(wellID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get samples"})
		return
	}
	fmt.Printf("Samples: %v\n", samples) // Log the samples
	c.JSON(http.StatusOK, samples)
}
