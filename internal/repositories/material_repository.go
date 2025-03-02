package repositories

import (
	"geo-app/internal/models"
	"log" // Импортируем пакет log

	"gorm.io/gorm"
)

type MaterialRepository interface {
	GetMaterialByID(id uint) (*models.Material, error)
	GetMaterialByName(name string) (*models.Material, error)
	GetAllMaterials() ([]*models.Material, error)
	CreateMaterial(material *models.Material) error
	UpdateMaterial(material *models.Material) error
	DeleteMaterial(id uint) error
}

type materialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db: db}
}

func (r *materialRepository) GetMaterialByID(id uint) (*models.Material, error) {
	var material models.Material
	result := r.db.Preload("Well").First(&material, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &material, nil
}

func (r *materialRepository) GetMaterialByName(name string) (*models.Material, error) {
	var material models.Material
	result := r.db.Preload("Well").Where("material_name = ?", name).First(&material)
	if result.Error != nil {
		return nil, result.Error
	}
	return &material, nil
}

func (r *materialRepository) GetAllMaterials() ([]*models.Material, error) {
	var materials []*models.Material
	result := r.db.Preload("Well").Find(&materials)
	if result.Error != nil {
		return nil, result.Error
	}
	return materials, nil
}

func (r *materialRepository) CreateMaterial(material *models.Material) error {
	result := r.db.Create(material)
	return result.Error
}

func (r *materialRepository) UpdateMaterial(material *models.Material) error {
	result := r.db.Save(material)
	return result.Error
}

func (r *materialRepository) DeleteMaterial(id uint) error {
	log.Printf("MaterialRepository: Deleting material with ID %d", id) // Добавили лог
	result := r.db.Delete(&models.Material{}, id)
	if result.Error != nil {
		log.Printf("MaterialRepository: Error deleting material %d: %v", id, result.Error) // Добавили лог
		return result.Error
	}
	log.Printf("MaterialRepository: Material %d deleted successfully", id) // Добавили лог
	return nil
}
