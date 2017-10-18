package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
