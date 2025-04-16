package models

import (
	"gorm.io/gorm"
)

// User моделі
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"` // Уникалды шектеу
	Password string `json:"password"`
	Role     string `json:"role"`
}
