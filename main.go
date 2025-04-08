package main

import (
	"fmt"
	"log"
	"my-clothing-store/config"
	"my-clothing-store/routes"
	"net/http"
)

func main() {
	config.InitDB()
	router := routes.SetupRoutes()
	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
