package main

import (
	"log"
	"os"
	"backend-picpay/internal/handlers"
	"backend-picpay/internal/models"
	"backend-picpay/internal/utils"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" 
	}
	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}