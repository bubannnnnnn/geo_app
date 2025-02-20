package repositories

import (
	"geo-app/internal/models"

	"gorm.io/gorm"
)

type SampleRepository struct {
	db *gorm.DB
}

func NewSampleRepository(db *gorm.DB) *SampleRepository {
	return &SampleRepository{db: db}
}

func (r *SampleRepository) Create(sample *models.Sample) error {
	return r.db.Create(sample).Error
}

func (r *SampleRepository) GetByID(id uint) (*models.Sample, error) {
	var sample models.Sample
	if err := r.db.First(&sample, id).Error; err != nil {
		return nil, err
	}
	return &sample, nil
}

func (r *SampleRepository) Update(sample *models.Sample) error {
	return r.db.Save(sample).Error
}

func (r *SampleRepository) Delete(id uint) error {
	return r.db.Delete(&models.Sample{}, id).Error
}

func (r *SampleRepository) GetSamplesByWellID(wellID uint) ([]models.Sample, error) { // Добавлено
	var samples []models.Sample
	if err := r.db.Where("well_id = ?", wellID).Find(&samples).Error; err != nil {
		return nil, err
	}
	return samples, nil
}
