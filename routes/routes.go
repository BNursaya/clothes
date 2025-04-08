package routes

import (
	"github.com/gorilla/mux"
	"my-clothing-store/controllers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/products", controllers.GetProducts).Methods("GET")
	return router
}
