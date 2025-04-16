package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Products    []Product `gorm:"foreignKey:CategoryID"` // Өнімдер мен категория арасындағы байланыс
}
