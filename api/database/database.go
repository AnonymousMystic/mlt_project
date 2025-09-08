package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"

	"github.com/google/uuid"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	mssdb *gorm.DB
	once  sync.Once
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

		mssdb = db

		log.Println("Connected to the database")
	})
}

// for authenticating trusted users
func FindUserWithUUID(uuid string) (string, error) {
	var results *string

	query := fmt.Sprintf(`
		SELECT uuid from Users
		WHERE uuid = '%s'
	`, uuid)

	mssdb.Raw(query).Scan(&results)

	// check if user is found
	if results != nil {
		return *results, nil
	}
	return "", errors.New("invalid Credentials")
}

// for authenticating untrusted users
func FindUserWithCredentials(email string, passwrd string) (string, error) {
	var results *string

	query := fmt.Sprintf(`
		SELECT uuid from Users
		WHERE email = '%s' AND passwrd = '%s'
	`, email, passwrd)

	mssdb.Raw(query).Scan(&results)

	// check if user is found
	if results != nil {
		return *results, nil
	}

	return "", errors.New("invalid Credentials")
}

// for authenticating newly registered users
func AddNewUser(email string, passwrd string) (string, error) {
	uuid := uuid.New().String()

	query := fmt.Sprintf(`
		INSERT INTO Users
		VALUES('%s', '%s', '%s', '%s')
	`, uuid, email, passwrd, "")

	status := mssdb.Exec(query)

	// handle query errors
	if status.Error != nil {
		log.Fatalf("Insert failed: %v", status.Error)
		return "", status.Error
	}

	return uuid, nil
}
