package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/roflomglol/gett/models"
)

// Handler is a struct for holding reference to db
type Handler struct {
	Db *gorm.DB
}

//GetDriver is a method for handling "/driver/:id" route
func (h *Handler) GetDriver(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	driver := h.Db.First(&models.Driver{}, id)

	if driver.RecordNotFound() {
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, driver.Value)
}

//BatchCreateDrivers is a method for handling "/import" route
func (h *Handler) BatchCreateDrivers(c echo.Context) error {
	raw, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return err
	}

	go func(body []byte) {
		var drivers []models.Driver

		if err := json.Unmarshal(body, &drivers); err != nil {
			c.Logger().Error(err)
			return
		}

		for _, driver := range drivers {
			if err := models.CreateDriver(&driver); err != nil {
				c.Logger().Error(err)
			}
		}
	}(raw)

	return c.NoContent(http.StatusAccepted)
}
