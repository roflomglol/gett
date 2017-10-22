package handlers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/roflomglol/gett/models"
	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo"
)

var (
	driverJSON = `{"id":2,"name":"Eyal Golan","license_number":"11-288-10"}`
	driver     = models.Driver{ID: 2, Name: "Eyal Golan", LicenseNumber: "11-288-10"}
	db         *gorm.DB
)

func testDB() *gorm.DB {
	var err error
	db, err = gorm.Open("postgres", "dbname=gett_test sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func TestGetDriver(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/driver/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	h := Handler{Db: testDB()}
	h.Db.Create(&driver)

	defer cleanAfterTestGetDriver(h.Db)

	// Assertions
	if assert.NoError(t, h.GetDriver(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, driverJSON, rec.Body.String())
	}
}

func TestGetDriverWhenNonexistent(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/driver/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/driver/:id")
	c.SetParamNames("id")
	c.SetParamValues("2")
	h := Handler{Db: testDB()}

	defer h.Db.Close()

	// Assertions
	if assert.NoError(t, h.GetDriver(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func cleanAfterTestGetDriver(db *gorm.DB) {
	db.Delete(&driver)
	db.Close()
}

func TestBatchCreateDrivers(t *testing.T) {
	// Setup
	e := echo.New()
	payload := strings.NewReader(`[{"id":2,"name":"Eyal Golan","license_number":"11-288-10"}]`)
	req := httptest.NewRequest(echo.POST, "/import", payload)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/import")
	h := Handler{Db: testDB()}

	defer h.Db.Close()

	// Assertions
	if assert.NoError(t, h.BatchCreateDrivers(c)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
	}
}

func TestBatchCreateDriversWithInvalidJson(t *testing.T) {
	// Setup
	e := echo.New()
	payload := strings.NewReader(`]`)
	req := httptest.NewRequest(echo.POST, "/import", payload)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/import")
	h := Handler{Db: testDB()}

	defer h.Db.Close()

	// Assertions
	if assert.NoError(t, h.BatchCreateDrivers(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
