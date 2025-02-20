package models

import (
	"gorm.io/gorm"
)

type Sample struct {
	gorm.Model
	WellID      uint    `json:"well_id"` // Foreign key to Well
	Name        string  `gorm:"not null" json:"name"`
	Description string  `json:"description"`
	Value       float64 `json:"value"` // Пример значения
}
