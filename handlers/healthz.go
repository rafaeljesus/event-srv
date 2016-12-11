package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HealthIndex(c echo.Context) error {
	response := map[string]string{"alive": "true"}

	return c.JSON(http.StatusOK, response)
}
