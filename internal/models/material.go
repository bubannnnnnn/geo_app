package models

import (
	"time"
)

type Material struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	MaterialName        string    `gorm:"not null;unique" json:"material_name"`
	MaterialDescription string    `json:"material_description"`
	PhysicalProperties  string    `json:"physical_properties"`
	ChemicalProperties  string    `json:"chemical_properties"`
	WellID              *uint     `json:"well_id"`                       // Foreign key for Well
	Well                Well      `gorm:"foreignKey:WellID" json:"well"` // связь с моделью Well
}
