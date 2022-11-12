package main

import (
	"fmt"
	"log"
	"os"
	"skyshi/controllers"
	"skyshi/routes"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DBNAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error DB")
	}

	// db.AutoMigrate(models.Activity{}, models.Todo{})

	skyshiService := controllers.SkyshiService{
		DB: db,
	}

	routes.Mux(skyshiService)

}
