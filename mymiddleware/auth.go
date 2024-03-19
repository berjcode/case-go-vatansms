package mymiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthenticationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return c.Redirect(http.StatusFound, "/login")
		}

		return next(c)
	}
}
