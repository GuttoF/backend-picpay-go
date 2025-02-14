package main

import (
	"log"
	"os"
	"backend-picpay/internal/handlers"
	"backend-picpay/internal/models"
	"backend-picpay/internal/utils"
	_ "backend-picpay/docs"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	"github.com/gofiber/fiber/v2"
)

// main is the entry point of the application. It connects to the database,
// performs database migrations, sets up the Fiber web server with routes,
// and starts the server on the specified port. If the PORT environment
// variable is not set, it defaults to port 3000. The application includes
// routes for creating users, transferring transactions, and serving Swagger
// documentation.
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


	// routes here
	app.Post("/users", handlers.CreateUser)
	app.Post("/transaction", handlers.Transfer)
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" 
	}
	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}