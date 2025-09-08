package database

import (
	"errors"
	"fmt"
	"golang-server/utils"
	"log"
	"os"
	"sync"
	"time"

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
	var userId *string
	if len(uuid) == 0 {
		return "", errors.New("no existing session")
	}

	query := fmt.Sprintf(`
		SELECT uuid from Users
		WHERE uuid = '%s'
	`, uuid)

	mssdb.Raw(query).Scan(&userId)

	// check if user is found
	if userId != nil {
		return *userId, nil
	}
	return "", errors.New("invalid Credentials")
}

// for authenticating untrusted users
func FindUserWithCredentials(email string, passwrd string) (utils.QueryUserIdResults, error) {
	if len(email) == 0 || len(passwrd) == 0 {
		return utils.QueryUserIdResults{}, errors.New("no existing session")
	}

	var results *utils.QueryUserIdResults

	query := fmt.Sprintf(`
		SELECT uuid, sessid from Users
		WHERE email = '%s' AND passwrd = '%s'
	`, email, passwrd)

	mssdb.Raw(query).Scan(&results)

	// check if user is found
	if results != nil {
		return *results, nil
	}

	return (utils.QueryUserIdResults{}), errors.New("invalid credentials")
}

// for checking for existing users
func FindUserWithUsername(email string) (bool, error) {
	var uuid *string
	if len(email) == 0 {
		return false, errors.New("not a valid input")
	}

	query := fmt.Sprintf(`
		SELECT uuid from Users
		WHERE email = '%s'
	`, email)

	mssdb.Raw(query).Scan(&uuid)

	// check if user is found
	if uuid == nil {
		return false, nil
	}

	return true, errors.New("invalid credentials")
}

// for authenticating newly registered users
func AddNewUser(email string, passwrd string) (string, string, error) {
	if len(email) == 0 || len(passwrd) == 0 {
		return "", "", errors.New("no existing session")
	}

	id := uuid.New().String()
	sessid := uuid.New().String()

	query := fmt.Sprintf(`
		INSERT INTO Users
		VALUES('%s', '%s', '%s', '%s')
	`, id, email, passwrd, sessid)

	status := mssdb.Exec(query)

	// handle query errors
	if status.Error != nil {
		log.Fatalf("Insert failed: %v", status.Error)
		return "", "", status.Error
	}

	return id, sessid, nil
}

// creates a session in the database
func CreateSession(sessid string, id string) error {
	if len(id) == 0 || len(sessid) == 0 {
		return errors.New("invalid id")
	}

	// generate session information
	issueDate := time.Now()
	sqlDateTime := issueDate.Format("2006-01-02 15:04:05")

	query := fmt.Sprintf(`
		INSERT INTO UserSessions
		VALUES('%s', '%s', '%s')
	`, sessid, id, sqlDateTime)

	status := mssdb.Exec(query)

	if status.Error != nil {
		log.Fatalf("Insert failed: %v", status.Error)
		return status.Error
	}

	return nil
}

// retrieves a session if it exists
func RetrieveSession(sessid string) (time.Time, error) {
	var sessionDate *time.Time

	if len(sessid) == 0 {
		return time.Time{}, errors.New("no existing session")
	}

	// try to find existing session
	query := fmt.Sprintf(`
		SELECT session_date FROM UserSessions
		WHERE sessionid = '%s'
	`, sessid)

	status := mssdb.Raw(query).Scan(&sessionDate)

	if status.Error != nil || sessionDate != nil {
		log.Fatalf("Insert failed: %v", status.Error)
		return time.Time{}, status.Error
	}

	return *sessionDate, nil
}

// access associated user using session information
func FindUserFromSession(sessid string) (string, error) {
	var uuid *string

	if len(sessid) == 0 {
		return "", errors.New("no existing session")
	}

	// try to find existing session
	query := fmt.Sprintf(`
		SELECT uuid FROM UserSessions
		WHERE sessionid = '%s'
	`, sessid)

	status := mssdb.Raw(query).Scan(&uuid)

	if status.Error != nil {
		log.Fatalf("Insert failed: %v", status.Error)
		return "", status.Error
	}

	return *uuid, nil
}

// invalidate and remove a session
func RemoveAndInvalidateSession(sessid string) error {
	var uuid *string
	if len(sessid) == 0 {
		return errors.New("no existing session")
	}

	removalQuery := fmt.Sprintf(`
		DELETE FROM UserSessions
		OUTPUT DELETED.uuid
		WHERE sessionid = '%s';
	`, sessid)

	removalStatus := mssdb.Raw(removalQuery).Scan(&uuid)

	// handle session removal query errors
	if removalStatus.Error != nil || uuid == nil {
		log.Fatalf("Removal failed: %v", removalStatus.Error)
		return removalStatus.Error
	}

	invalidationQuery := fmt.Sprintf(`
		UPDATE Users
		SET sessid = NULL
		WHERE uuid = '%s';
	`, *uuid)

	invalidationStatus := mssdb.Exec(invalidationQuery)

	// handle session invalidation query errors
	if invalidationStatus.Error != nil {
		log.Fatalf("Invalidation failed: %v", invalidationStatus.Error)
		return removalStatus.Error
	}

	return nil
}
