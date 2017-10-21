package database

import (
	"errors"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Db is a var holding reference to a database connection
var Db *gorm.DB

func init() {
	var err error

	databaseURL, err := getDatabaseURL()

	if err != nil {
		log.Fatal(err)
	}

	Db, err = gorm.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}
}

func getDatabaseURL() (string, error) {
	databaseURL, exists := os.LookupEnv("DATABASE_URL")

	if !exists {
		return databaseURL, errors.New("DATABASE_URL env varibale is not set")
	}

	return databaseURL, nil
}
