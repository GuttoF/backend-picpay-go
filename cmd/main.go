package main

import (
	"backend-picpay/internal/models"
	"backend-picpay/internal/utils"

	"log"
)

func main() {
	db := utils.ConnectDB()
	if db == nil {
		log.Fatalf("Failed to connect to database")
	}

	err := db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transaction{}, &models.Notification{})
	if err != nil {
		log.Fatalf("Failed to migrate database:  %v", err)
	}

	log.Println("Database migration completed succesfully")
}