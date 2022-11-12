package main

import (
	"fmt"
	"log"
	"os"
	"skyshi/controllers"
	"skyshi/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("mysql://%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_HOST"),
			os.Getenv("MYSQL_DBNAME")))
	if err != nil {
		log.Fatal("ERROR CONNECT", err)
	}
	err = m.Up()
	if err != nil {
		log.Fatal("ERROR Migrate", err)
	} else {
		log.Println("SUCCESS Migrate")
	}
}

func main() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DBNAME"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error DB")
	}

	skyshiService := controllers.SkyshiService{
		DB: db,
	}

	routes.Mux(skyshiService)

}
