package utils

import (
    "os"
    "log"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
    "github.com/joho/godotenv"
)

func ConnectDB() *gorm.DB {
    err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file:", err)
        return nil
    }

    dbHost := os.Getenv("POSTGRES_HOST")
    dbPort := os.Getenv("POSTGRES_PORT")
    dbUser := os.Getenv("POSTGRES_USER")
    dbPassword := os.Getenv("POSTGRES_PASSWORD")
    dbName := os.Getenv("POSTGRES_DB")
    sslMode := os.Getenv("POSTGRES_SSLMODE")

    dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser +
        " password=" + dbPassword + " dbname=" + dbName + " sslmode=" + sslMode

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Println("Error connecting to database:", err)
        return nil
    }

    return db
}
