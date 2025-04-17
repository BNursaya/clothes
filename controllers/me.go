package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func GetProfile(c *gin.Context) {
	userToken, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	claims := userToken.(*jwt.Token).Claims.(jwt.MapClaims)

	c.JSON(http.StatusOK, gin.H{
		"id":    claims["id"],
		"email": claims["email"],
		"role":  claims["role"],
	})
}
