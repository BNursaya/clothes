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

//package main
//
//import (
//	"fmt"
//	"golang.org/x/crypto/bcrypt"
//)
//
//func main() {
//	hash, _ := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
//	fmt.Println(string(hash))
//}
