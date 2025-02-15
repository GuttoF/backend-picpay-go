package utils

import (
    "fmt"
    "log"
    "os"
    "sync"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var (
    DB   *gorm.DB
    once sync.Once
)

func ConnectDB() {
    once.Do(func() {
        if err := godotenv.Load(); err != nil {
            log.Println("Warning: .env file not found, using system environment variables")
        }

        dbHost := os.Getenv("POSTGRES_HOST")
        dbPort := os.Getenv("POSTGRES_PORT")
        dbUser := os.Getenv("POSTGRES_USER")
        dbPassword := os.Getenv("POSTGRES_PASSWORD")
        dbName := os.Getenv("POSTGRES_DB")
        sslMode := os.Getenv("POSTGRES_SSLMODE")

        dsn := fmt.Sprintf(
            "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
            dbHost, dbPort, dbUser, dbPassword, dbName, sslMode,
        )

        var err error
        DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            log.Fatal("Error connecting to database:", err)
        }

        log.Println("Database connection established successfully")
    })
}

func GetDB() *gorm.DB {
    if DB == nil {
        log.Fatal("Database is not initialized. Call ConnectDB() first.")
    }
    return DB
}
