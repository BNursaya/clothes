package controllers

import (
	"github.com/gin-gonic/gin"
	"myProject/config"
	"myProject/models"
	"net/http"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating order"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	if err := config.DB.Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching orders"})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}
	c.JSON(http.StatusOK, order)
}

func UpdateOrderStatus(c *gin.Context) {
	id := c.Param("id")
	var update map[string]interface{}
	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := config.DB.Model(&models.Order{}).Where("id = ?", id).Updates(update).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating order"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Order updated"})
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Order{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting order"})
		return
	}
	c.Status(http.StatusNoContent)
}
