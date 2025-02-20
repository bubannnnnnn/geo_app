package services

import (
	"geo-app/internal/models"
	"geo-app/internal/repositories"
)

type SampleService struct {
	sampleRepo *repositories.SampleRepository
}

func NewSampleService(sampleRepo *repositories.SampleRepository) *SampleService {
	return &SampleService{sampleRepo: sampleRepo}
}

func (s *SampleService) CreateSample(sample *models.Sample) error {
	return s.sampleRepo.Create(sample)
}

func (s *SampleService) GetSampleByID(id uint) (*models.Sample, error) {
	return s.sampleRepo.GetByID(id)
}

func (s *SampleService) UpdateSample(sample *models.Sample) error {
	return s.sampleRepo.Update(sample)
}

func (s *SampleService) DeleteSample(id uint) error {
	return s.sampleRepo.Delete(id)
}

func (s *SampleService) GetSamplesByWellID(wellID uint) ([]models.Sample, error) { // Добавлено
	return s.sampleRepo.GetSamplesByWellID(wellID)
}
