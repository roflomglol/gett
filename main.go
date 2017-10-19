package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
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
	LicenseNumber string `gorm:"type:varchar(255);not null;unique" json:"license_number"`
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
	e.GET("driver/:id", getDriver)
	e.POST("import", batchCreateDrivers)

	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", serverPort)))
}

func getDriver(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	driver := db.First(&Driver{}, id)

	if driver.RecordNotFound() {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, driver.Value)
}

func batchCreateDrivers(c echo.Context) error {
	raw, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	go func(body []byte) {
		var drivers []Driver

		if err := json.Unmarshal(body, &drivers); err != nil {
			c.Logger().Error(err)
		}

		for _, driver := range drivers {
			if err := CreateDriver(&driver); err != nil {
				c.Logger().Error(err)
			}
		}
	}(raw)

	return c.NoContent(http.StatusAccepted)
}

func CreateDriver(d *Driver) error {
	return db.Create(d).Error
}

func (d *Driver) BeforeSave() error {
	return d.Validate()
}

func (d *Driver) Validate() error {
	if d.Exists(d.ID) {
		return fmt.Errorf("driver with ID %v already exists", d.ID)
	}

	if d.Name == "" {
		return errors.New("driver's name can't be blank")
	}

	if d.LicenseNumber == "" {
		return errors.New("driver's license number can't be blank")
	}

	if isLicenseNumberExists(d.LicenseNumber) {
		return fmt.Errorf("driver with license number %v already exists", d.LicenseNumber)
	}

	return nil
}

func (d *Driver) Exists(id uint) bool {
	return !db.Where("id = ?", id).First(&Driver{}).RecordNotFound()
}

func isLicenseNumberExists(l string) bool {
	return !db.Where("license_number = ?", l).First(&Driver{}).RecordNotFound()
}
