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
	utils.ConnectDB()
	db := utils.GetDB()

	if err := db.AutoMigrate(&models.User{}, &models.Wallet{}, &models.Transfer{}, &models.Notification{}); err != nil {
		log.Fatalf("Database migration failed: %v", err)
	}
	log.Println("Database migration completed successfully")

	app := fiber.New()

	//API routes
	log.Println("Setting up API routes...")
	app.Post("/users", handlers.CreateUserHandler(db))
	app.Post("/transfers", handlers.TransferHandler(db))
	app.Post("/wallets", handlers.CreateWalletHandler(db))
	app.Post("/wallets/deposit", handlers.DepositHandler(db))
	app.Post("/wallets/withdraw", handlers.WithdrawHandler(db))
	app.Get("/wallets/:user_id", handlers.GetWalletByUserIDHandler(db))

	// Swagger documentation
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("Server is running on port %s", port)
	log.Fatal(app.Listen(":" + port))
}
