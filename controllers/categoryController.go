package controllers

import (
	"github.com/gin-gonic/gin"
	"myProject/config"
	"myProject/models"
	"net/http"
)

// Категорияларды алу
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := config.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// Жаңа категория қосу
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := config.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating category"})
		return
	}

	c.JSON(http.StatusOK, category)
}
