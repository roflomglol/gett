package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
)

type Driver struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Name          string `gorm:"type:varchar(255);not null;" json:"name"`
	LicenseNumber string `gorm:"type:varchar(255);not null;unique" json:"licenseNumber"`
}

var db = initDb()

func initDb() *gorm.DB {
	databaseURL := os.Getenv("DATABASE_URL")
	db, err := gorm.Open("postgres", databaseURL)

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Driver{})

	return db
}

func main() {
	e := echo.New()

	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", serverPort)))
}

