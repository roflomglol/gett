package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/roflomglol/gett/database"
)

// Driver struct represents driver model
type Driver struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`

	Name          string `gorm:"type:varchar(255);not null;" json:"name"`
	LicenseNumber string `gorm:"type:varchar(255);not null;unique" json:"license_number"`
}

func init() {
	database.Db.AutoMigrate(&Driver{})
}

// CreateDriver saves driver to database
func CreateDriver(d *Driver) error {
	return database.Db.Create(d).Error
}

// BeforeSave is a gorm callback method
func (d *Driver) BeforeSave() error {
	return d.validate()
}

func (d *Driver) validate() error {
	if d.exists(d.ID) {
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

func (d *Driver) exists(id uint) bool {
	return !database.Db.Where("id = ?", id).First(&Driver{}).RecordNotFound()
}

func isLicenseNumberExists(l string) bool {
	return !database.Db.Where("license_number = ?", l).First(&Driver{}).RecordNotFound()
}
