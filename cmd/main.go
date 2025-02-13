package main

import (
	"backend-picpay/internal/models"
	"backend-picpay/internal/utils"
	"backend-picpay/internal/handlers"
	"log"
	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()

	app.Post("/users", handlers.CreateUser)
	app.Post("/transactions", handlers.Transfer)

	log.Fatal(app.Listen(":3000"))
}