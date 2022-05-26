package utils

import (
	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, status int, err error) error {
	data := map[string]interface{}{
		"status": "failure",
		"reason": err.Error(),
	}

	return c.JSON(status, data)
}

func ErrorResponsef(c echo.Context, status int, err string) error {
	data := map[string]interface{}{
		"status": "failure",
		"reason": err,
	}

	return c.JSON(status, data)
}
