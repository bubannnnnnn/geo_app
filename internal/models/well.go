package models

import (
	"gorm.io/gorm"
)

type Well struct {
	gorm.Model
	Name        string  `gorm:"not null" json:"name"`
	Status      string  `json:"status"` //"active", "inactive", ...
	Latitude    float64 `gorm:"not null" json:"latitude"`
	Longitude   float64 `gorm:"not null" json:"longitude"`
	Description string  `json:"description"`

	ResourceType string  `json:"resourceType"` // Тип ископаемого (например, нефть, газ, уголь)
	Depth        float64 `json:"depth"`        // Глубина залегания (в метрах)
	Volume       float64 `json:"volume"`       // Примерный объем (в тоннах, баррелях и т.д.)
	GroundType   string  `json:"groundType"`   // Тип грунта (например, песчаник, глина)

	CadastralNumber string `json:"cadastralNumber"` // Кадастровый номер
	ImageURL        string `json:"imageURL"`        // URL картинки
}
