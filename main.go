package main

import (
	"myProject/config"
	"myProject/routes"
)

func main() {
	config.InitDB()
	router := routes.SetupRoutes()
	router.Run(":8080")
}
