package routes

import (
	"github.com/gin-gonic/gin"
	"myProject/controllers"
	"myProject/middleware"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	auth := router.Group("/api")
	auth.Use(middleware.TokenAuthMiddleware())
	{
		auth.GET("/me", controllers.GetProfile)

		auth.GET("/products", controllers.GetProducts)
		auth.GET("/products/:id", controllers.GetProductByID)
		auth.POST("/products", controllers.CreateProduct)
		auth.PUT("/products/:id", controllers.UpdateProduct)
		auth.DELETE("/products/:id", controllers.DeleteProduct)
		auth.GET("/products/search", controllers.SearchProducts)

		auth.GET("/categories", controllers.GetCategories)
		auth.POST("/categories", controllers.CreateCategory)

		auth.POST("/orders", controllers.CreateOrder)
		auth.GET("/orders", controllers.GetOrders)
		auth.GET("/orders/:id", controllers.GetOrderById)
		auth.PUT("/orders/:id", controllers.UpdateOrderStatus)
		auth.DELETE("/orders/:id", controllers.DeleteOrder)
	}

	return router
}
