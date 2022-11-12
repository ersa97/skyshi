package main

import (
	"log"
	"skyshi/controllers"
	"skyshi/database"
	"skyshi/routes"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db := database.ConnectMariaDB(3) // init database connection

	skyshiService := controllers.SkyshiService{
		DB: db,
	}

	routes.Mux(skyshiService)

}
