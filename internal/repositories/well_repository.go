package repositories

import (
	"geo-app/internal/models"

	"gorm.io/gorm"
)

type WellRepository interface {
	Create(well *models.Well) error
	GetByID(id uint) (*models.Well, error)
	GetAll() ([]models.Well, error)
	Update(well *models.Well) error
	Delete(id uint) error
}

type wellRepository struct {
	db *gorm.DB
}

func NewWellRepository(db *gorm.DB) WellRepository {
	return &wellRepository{db: db}
}

func (r *wellRepository) Create(well *models.Well) error {
	return r.db.Create(well).Error
}

func (r *wellRepository) GetByID(id uint) (*models.Well, error) {
	var well models.Well
	if err := r.db.First(&well, id).Error; err != nil {
		return nil, err
	}
	return &well, nil
}

func (r *wellRepository) GetAll() ([]models.Well, error) {
	var wells []models.Well
	if err := r.db.Find(&wells).Error; err != nil {
		return nil, err
	}
	return wells, nil
}

func (r *wellRepository) Update(well *models.Well) error {
	return r.db.Save(well).Error
}

func (r *wellRepository) Delete(id uint) error {
	return r.db.Delete(&models.Well{}, id).Error
}
