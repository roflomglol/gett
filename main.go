package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/roflomglol/gett/database"
	"github.com/roflomglol/gett/models"
)

func main() {
	defer database.Db.Close()

	e := echo.New()
	e.GET("driver/:id", getDriver)
	e.POST("import", batchCreateDrivers)

	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", serverPort)))
}

func getDriver(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	driver := database.Db.First(&models.Driver{}, id)

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
