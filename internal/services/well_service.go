package services

import (
	"geo-app/internal/models"
	"geo-app/internal/repositories"
)

type WellService interface {
	GetWells() ([]models.Well, error)
	GetWellByID(id uint) (*models.Well, error)
	CreateWell(well models.Well) (*models.Well, error)
	UpdateWell(well models.Well) (*models.Well, error)
	DeleteWell(id uint) error
}

type wellService struct {
	wellRepository repositories.WellRepository
}

func NewWellService(wellRepository repositories.WellRepository) WellService {
	return &wellService{wellRepository: wellRepository}
}

func (s *wellService) GetWells() ([]models.Well, error) {
	return s.wellRepository.GetAll()
}

func (s *wellService) GetWellByID(id uint) (*models.Well, error) {
	return s.wellRepository.GetByID(id)
}

func (s *wellService) CreateWell(well models.Well) (*models.Well, error) {
	err := s.wellRepository.Create(&well)
	if err != nil {
		return nil, err
	}
	return &well, nil
}

func (s *wellService) UpdateWell(well models.Well) (*models.Well, error) {
	err := s.wellRepository.Update(&well)
	if err != nil {
		return nil, err
	}
	return &well, nil
}

func (s *wellService) DeleteWell(id uint) error {
	return s.wellRepository.Delete(id)
}
