package models

import (
	"gorm.io/gorm"
)

type Well struct {
	gorm.Model
	Name        string  `json:"name" binding:"required"`
	Latitude    float64 `json:"latitude" binding:"required,latitude"`
	Longitude   float64 `json:"longitude" binding:"required,longitude"`
	Depth       float64 `json:"depth"`
	Description string  `json:"description"`
	// Добавьте другие поля, которые вам нужны
}
