package main

import (
	"fmt"
	"os"

	"github.com/roflomglol/gett/handlers"

	"github.com/labstack/echo"
	"github.com/roflomglol/gett/database"
)

func main() {
	defer database.Db.Close()

	e := echo.New()
	h := handlers.Handler{Db: database.Db}
	e.GET("/driver/:id", h.GetDriver)
	e.POST("/import", h.BatchCreateDrivers)

	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", serverPort)))
}
