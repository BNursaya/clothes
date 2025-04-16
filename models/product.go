package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// Product моделі
type Product struct {
	gorm.Model
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	CategoryID  int            `json:"category_id"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
	ImageURL    string         `json:"image_url"`
	Sizes       datatypes.JSON `json:"sizes"`
	Colors      datatypes.JSON `json:"colors"`
	Material    string         `json:"material"`
	Gender      string         `json:"gender"`
}
