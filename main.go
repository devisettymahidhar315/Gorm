package main

import (
	"app/adapaters"
	"app/api"
	"app/usecase"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database connection string
	dsn := "root:rootpassword@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	database_init := adapaters.Init(db)
	usecae_init := usecase.Init(database_init)
	api_init := api.Init(usecae_init)
	api_init.Post()
}
