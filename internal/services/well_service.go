package services

import (
	"geo-app/internal/models"
	"geo-app/internal/repositories"
)

type WellService interface {
	GetWells() ([]models.Well, error)
	GetWellByID(id uint) (models.Well, error)
	CreateWell(well models.Well) (models.Well, error)
	UpdateWell(well models.Well) (models.Well, error)
	DeleteWell(id uint) error
}

type wellService struct {
	wellRepository repositories.WellRepository
}

func NewWellService(wellRepository repositories.WellRepository) WellService {
	return &wellService{wellRepository: wellRepository}
}

func (s *wellService) GetWells() ([]models.Well, error) {
	return s.wellRepository.GetWells()
}

func (s *wellService) GetWellByID(id uint) (models.Well, error) {
	return s.wellRepository.GetWellByID(id)
}

func (s *wellService) CreateWell(well models.Well) (models.Well, error) {
	return s.wellRepository.CreateWell(well)
}

func (s *wellService) UpdateWell(well models.Well) (models.Well, error) {
	return s.wellRepository.UpdateWell(well)
}

func (s *wellService) DeleteWell(id uint) error {
	return s.wellRepository.DeleteWell(id)
}
