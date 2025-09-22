package main

import (
	"log"
	
	"github.com/joho/godotenv"
	"github.com/Zhaobo-Wang/go-projects/database"
	"github.com/Zhaobo-Wang/go-projects/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found")
	}
	
	database.ConnectDatabase()
	
	r := routes.SetupRouter()
	
	r.Run(":8080")
}