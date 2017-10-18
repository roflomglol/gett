package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Driver struct {
	gorm.Model

	Name          string `gorm:"type:varchar(255);not null;unique"`
	LicenseNumber string `gorm:"type:varchar(255);not null;unique"`
}

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Driver{})

	defer db.Close()
}
