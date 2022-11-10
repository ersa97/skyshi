package main

import (
	"log"
	"os"
	"skyshi/controllers"
	"skyshi/database"
	"skyshi/routes"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	APPLICATION_PORT string
	DATABASE_URL     string
	db               *gorm.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	APPLICATION_PORT = os.Getenv("APPLICATION_PORT")
	DATABASE_URL = os.Getenv("DATABASE_URL")

	db := database.ConnectMariaDB(3) // init database connection

	skyshiService := controllers.SkyshiService{
		DB: db,
	}

	routes.Mux(skyshiService)

}
