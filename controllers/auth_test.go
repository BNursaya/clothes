package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"myProject/config"
	"myProject/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

var testUserID uint
var testCategoryID uint

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	r.POST("/register", Register)
	r.POST("/login", Login)

	r.POST("/api/products", CreateProduct)
	r.GET("/api/products", GetProducts)
	r.GET("/api/products/:id", GetProductByID)
	r.PUT("/api/products/:id", UpdateProduct)
	r.DELETE("/api/products/:id", DeleteProduct)

	r.POST("/api/orders", CreateOrder)
	r.GET("/api/orders", GetOrders)

	return r
}

func initTestDB() {
	config.InitDB()
	config.DB.Exec("DELETE FROM orders")
	config.DB.Exec("DELETE FROM products")
	config.DB.Exec("DELETE FROM categories")
	config.DB.Exec("DELETE FROM users")

	// Категория қосу
	config.DB.Raw(`INSERT INTO categories (name, description, created_at, updated_at)
		VALUES (?, ?, now(), now()) RETURNING id`, "Test Category", "For testing").Scan(&testCategoryID)

	// Юзер қосу
	hashedPassword := "$2a$10$exQk6/eEoZMbd1qxnAJOo.PKT4DC5f4Y9xRk.hnFgQ79RGsTwOW8C"
	config.DB.Raw(`INSERT INTO users (name, email, password, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, now(), now()) RETURNING id`,
		"Test User", "test@example.com", hashedPassword, "user").Scan(&testUserID)

	fmt.Println("Test category ID:", testCategoryID)
	fmt.Println("Test user ID:", testUserID)
}

func TestRegister(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	body := map[string]string{
		"name":     "TestUser2",
		"email":    "test2@example.com",
		"password": "test123",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestLogin(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	body := map[string]string{
		"email":    "test@example.com",
		"password": "123456",
	}
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateProduct(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	product := models.Product{Name: "TestProduct", Price: 1000, CategoryID: int(testCategoryID)}
	body, _ := json.Marshal(product)
	req, _ := http.NewRequest("POST", "/api/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetProducts(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	config.DB.Create(&models.Product{Name: "Shirt", Price: 2000, CategoryID: int(testCategoryID)})
	req, _ := http.NewRequest("GET", "/api/products", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetProductByID(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	product := models.Product{Name: "SingleProduct", Price: 1500, CategoryID: int(testCategoryID)}
	config.DB.Create(&product)
	req, _ := http.NewRequest("GET", "/api/products/"+strconv.Itoa(int(product.ID)), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestCreateOrder(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	product := models.Product{Name: "OrderProduct", Price: 1200, CategoryID: int(testCategoryID)}
	config.DB.Create(&product)
	order := models.Order{UserID: testUserID, ProductID: product.ID, Quantity: 2, Status: "pending"}
	body, _ := json.Marshal(order)
	req, _ := http.NewRequest("POST", "/api/orders", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetOrders(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	product := models.Product{Name: "OrderProduct2", Price: 1400, CategoryID: int(testCategoryID)}
	config.DB.Create(&product)
	config.DB.Create(&models.Order{UserID: testUserID, ProductID: product.ID, Quantity: 1, Status: "confirmed"})
	req, _ := http.NewRequest("GET", "/api/orders", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestGetProductsWithFilter(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	config.DB.Create(&models.Product{Name: "FilterProduct", Price: 2500, CategoryID: int(testCategoryID)})
	req, _ := http.NewRequest("GET", "/api/products?min_price=2000&max_price=3000", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestUpdateProduct(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	product := models.Product{Name: "UpdateMe", Price: 1000, CategoryID: int(testCategoryID)}
	config.DB.Create(&product)
	update := map[string]interface{}{
		"name":  "UpdatedName",
		"price": 2000,
	}
	body, _ := json.Marshal(update)
	req, _ := http.NewRequest("PUT", "/api/products/"+strconv.Itoa(int(product.ID)), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
}

func TestDeleteProduct(t *testing.T) {
	initTestDB()
	r := setupTestRouter()
	product := models.Product{Name: "DeleteMe", Price: 3000, CategoryID: int(testCategoryID)}
	config.DB.Create(&product)
	req, _ := http.NewRequest("DELETE", "/api/products/"+strconv.Itoa(int(product.ID)), nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNoContent, w.Code)
}
