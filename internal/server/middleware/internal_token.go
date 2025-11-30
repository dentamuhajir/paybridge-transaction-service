package middleware

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ValidateInternalToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")

		expected := "Bearer " + os.Getenv("TOKEN_INTERNAL_SERVICE")

		if auth != expected {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid internal token",
			})
		}

		return next(c)
	}
}
