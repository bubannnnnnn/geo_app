package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`    // Хранить только хеш пароля
	Email    string `gorm:"uniqueIndex"` // Добавляем поле email
}
