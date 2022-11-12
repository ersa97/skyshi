package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var sqlDb *sql.DB
var gormDb *gorm.DB
var ModelsDb *gorm.DB

func ConnectMariaDB(retries int) *gorm.DB {

	if retries > 1 {
		log.Printf("Retrying connect to DB instance, Attempt %v", strconv.Itoa(retries))

		if retries > 5 {
			log.Printf("Cannot recovery situation retries > 5 attempt")
			os.Exit(1)
		}
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DBNAME"))

	sqlDB, err := sql.Open("mysql", connString)

	if err != nil {
		log.Printf("error on creating connection sql database %v", err)
		ConnectMariaDB(retries + 1)
		return nil
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	// gormDB, err := gorm.Open("mysql", sqlDB)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Println("error on creating gorm connection ", err)
		ConnectMariaDB(retries + 1)
		return nil
	}

	if err != nil {
		log.Printf("error on creating connection database %v", err)

		ConnectMariaDB(retries + 1)
		return nil
	}

	newLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
		})
	gormDB.Session(&gorm.Session{Logger: newLogger})

	log.Println("database connection successfully")

	return gormDB
}
