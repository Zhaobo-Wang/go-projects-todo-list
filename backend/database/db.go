package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Zhaobo-Wang/go-projects/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	if username == "" {
		username = "root"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3306"
	}
	if dbname == "" {
		dbname = "todo_app"
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		username, password, host, port, dbname)
	
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB.Exec("CREATE DATABASE IF NOT EXISTS " + dbname)
	
	if err := DB.AutoMigrate(&models.User{}); err != nil {
		log.Printf("Error migrating User model: %v", err)
	}
	
	if err := DB.AutoMigrate(&models.Todo{}); err != nil {
		log.Printf("Error migrating Todo model: %v", err)
	}
}