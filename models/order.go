package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	ProductID uint   `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Status    string `json:"status"`
}
