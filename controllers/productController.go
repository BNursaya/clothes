package controllers

import (
	"github.com/gin-gonic/gin"
	"myProject/config"
	"myProject/models"
	"net/http"
	"strconv"
)

// Жаңа өнім қосу
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if product.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be positive"})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Өнімдерді фильтрациямен алу
func GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit

	category := c.Query("category")
	minPrice := c.Query("min_price")
	maxPrice := c.Query("max_price")

	var products []models.Product
	query := config.DB.Preload("Category") // Категорияны жүктеу

	// Категория фильтрі
	if category != "" {
		query = query.Where("category_id = ?", category)
	}

	// Баға фильтрі
	if minPrice != "" && maxPrice != "" {
		min, _ := strconv.ParseFloat(minPrice, 64)
		max, _ := strconv.ParseFloat(maxPrice, 64)
		query = query.Where("price BETWEEN ? AND ?", min, max)
	}

	err := query.Offset(offset).Limit(limit).Find(&products).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching products: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

// ID арқылы өнімді алу
func GetProductByID(c *gin.Context) {
	id := c.Param("id")

	var product models.Product
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Өнімді жаңарту
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct models.Product

	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data"})
		return
	}

	if updatedProduct.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Price must be positive"})
		return
	}

	if err := config.DB.Model(&models.Product{}).Where("id = ?", id).Updates(updatedProduct).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating product"})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

// Өнімді жою
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	if err := config.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting product"})
		return
	}

	c.Status(http.StatusNoContent)
}

// Іздеу (поиск)
func SearchProducts(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query missing"})
		return
	}

	var products []models.Product
	if err := config.DB.Where("name ILIKE ?", "%"+query+"%").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error searching products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
