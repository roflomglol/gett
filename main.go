package main

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Driver struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Name          string `gorm:"type:varchar(255);not null;" json:"name"`
	LicenseNumber string `gorm:"type:varchar(255);not null;unique" json:"licenseNumber"`
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
