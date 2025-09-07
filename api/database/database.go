package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"golang-server/models"

	"github.com/joho/godotenv"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	once sync.Once
)

func ConnectDatabase() {
	once.Do(func() {
		err := godotenv.Load() // Load .env file
		if err != nil {
			log.Fatalf("Error loading .env file")
		}

		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
			os.Getenv("SQL_SERVER_USER"),
			os.Getenv("SQL_SERVER_PASS"),
			os.Getenv("SQL_SERVER_HOST"),
			os.Getenv("SQL_SERVER_PORT"),
			os.Getenv("SQL_SERVER_DB"),
		)

		db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database: ", err)
		}

		DB = db

		log.Println("Connected to the database")
		db.AutoMigrate(&models.User{})
	})
}
