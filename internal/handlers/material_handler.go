package handlers

import (
	"geo-app/internal/models"
	"geo-app/internal/services"
	"log" // Импортируем пакет log
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MaterialHandler struct {
	materialService services.MaterialService
}

func NewMaterialHandler(materialService services.MaterialService) *MaterialHandler {
	return &MaterialHandler{materialService: materialService}
}

func (h *MaterialHandler) GetMaterialByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	material, err := h.materialService.GetMaterialByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *MaterialHandler) GetMaterialByName(c *gin.Context) {
	materialName := c.Param("material_name")

	material, err := h.materialService.GetMaterialByName(materialName)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Material not found"})
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *MaterialHandler) GetAllMaterials(c *gin.Context) {
	materials, err := h.materialService.GetAllMaterials()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get materials"})
		return
	}
	c.JSON(http.StatusOK, materials)
}

func (h *MaterialHandler) CreateMaterial(c *gin.Context) {
	var material models.Material
	if err := c.ShouldBindJSON(&material); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.materialService.CreateMaterial(&material)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create material"})
		return
	}

	c.JSON(http.StatusCreated, material)
}

func (h *MaterialHandler) UpdateMaterial(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var material models.Material
	if err = c.ShouldBindJSON(&material); err != nil { //  Изменили := на =
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	material.ID = uint(id)

	err = h.materialService.UpdateMaterial(&material)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update material"})
		return
	}

	c.JSON(http.StatusOK, material)
}

func (h *MaterialHandler) DeleteMaterial(c *gin.Context) {
	log.Println("DeleteMaterial handler called") // Добавили лог
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("Invalid ID") // Добавили лог
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	log.Printf("Deleting material with ID: %d", id) // Добавили лог
	err = h.materialService.DeleteMaterial(uint(id))
	if err != nil {
		log.Println("Error deleting material:", err) // Добавили лог
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete material"})
		return
	}

	log.Println("Material deleted successfully") // Добавили лог
	c.Status(http.StatusNoContent)
}
