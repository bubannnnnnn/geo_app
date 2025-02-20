// internal/repositories/well_repository.go
package repositories

import (
	"geo-app/internal/models"

	"gorm.io/gorm"
)

type WellRepository interface {
	GetWells() ([]models.Well, error)
	GetWellByID(id uint) (models.Well, error)
	CreateWell(well models.Well) (models.Well, error)
	UpdateWell(well models.Well) (models.Well, error)
	DeleteWell(id uint) error
}

type wellRepository struct {
	db *gorm.DB
}

func NewWellRepository(db *gorm.DB) WellRepository {
	return &wellRepository{db: db}
}

func (r *wellRepository) GetWells() ([]models.Well, error) {
	var wells []models.Well
	result := r.db.Find(&wells)
	return wells, result.Error
}

func (r *wellRepository) GetWellByID(id uint) (models.Well, error) {
	var well models.Well
	result := r.db.First(&well, id)
	return well, result.Error
}

func (r *wellRepository) CreateWell(well models.Well) (models.Well, error) {
	result := r.db.Create(&well)
	return well, result.Error
}

func (r *wellRepository) UpdateWell(well models.Well) (models.Well, error) {
	result := r.db.Model(&well).Updates(well)
	return well, result.Error
}

func (r *wellRepository) DeleteWell(id uint) error {
	result := r.db.Delete(&models.Well{}, id)
	return result.Error
}
