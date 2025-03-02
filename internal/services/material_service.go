package services

import (
	"geo-app/internal/models"
	"geo-app/internal/repositories"
	"log" // Импортируем пакет log
)

type MaterialService interface {
	GetMaterialByID(id uint) (*models.Material, error)
	GetMaterialByName(name string) (*models.Material, error)
	GetAllMaterials() ([]*models.Material, error)
	CreateMaterial(material *models.Material) error
	UpdateMaterial(material *models.Material) error
	DeleteMaterial(id uint) error
}

type materialService struct {
	repo repositories.MaterialRepository
}

func NewMaterialService(repo repositories.MaterialRepository) MaterialService {
	return &materialService{repo: repo}
}

func (s *materialService) GetMaterialByID(id uint) (*models.Material, error) {
	return s.repo.GetMaterialByID(id)
}

func (s *materialService) GetMaterialByName(name string) (*models.Material, error) {
	return s.repo.GetMaterialByName(name)
}

func (s *materialService) GetAllMaterials() ([]*models.Material, error) {
	return s.repo.GetAllMaterials()
}

func (s *materialService) CreateMaterial(material *models.Material) error {
	return s.repo.CreateMaterial(material)
}

func (s *materialService) UpdateMaterial(material *models.Material) error {
	return s.repo.UpdateMaterial(material)
}

func (s *materialService) DeleteMaterial(id uint) error {
	log.Printf("MaterialService: Deleting material with ID %d", id) // Добавили лог
	err := s.repo.DeleteMaterial(id)
	if err != nil {
		log.Printf("MaterialService: Error deleting material %d: %v", id, err) // Добавили лог
		return err
	}
	log.Printf("MaterialService: Material %d deleted successfully", id) // Добавили лог
	return nil
}
