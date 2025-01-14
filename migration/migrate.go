package main

import (
	"log"

	"myproject/database"
	"myproject/models"
)

func Migrate() {
	db := database.GetDb()
	db.AutoMigrate(&models.User{}, &models.Album{})
}

func main() {
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	Migrate()
}
