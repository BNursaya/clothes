package controllers

import (
	"encoding/json"
	"fmt"
	"my-clothing-store/config"
	"my-clothing-store/models"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query("SELECT id, name, description, price, category_id, image_url FROM products")
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.CategoryID, &product.ImageURL)
		if err != nil {
			http.Error(w, "Error scanning products", http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}
