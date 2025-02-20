package services

import (
	"geo-app/internal/models"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	db *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{db: db}
}

func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *AuthService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *AuthService) CreateUser(user *models.User) error {
	hashedPassword, err := s.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	result := s.db.Create(user)
	if result.Error != nil {
		log.Println("Failed to create user:", result.Error) // логируем ошибку
		return result.Error
	}
	return nil
}

func (s *AuthService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := s.db.Where("username = ?", username).First(&user)
	if result.Error != nil {
		log.Println("Failed to get user by username:", result.Error)
		return nil, result.Error
	}
	return &user, nil
}
