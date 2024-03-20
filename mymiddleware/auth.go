package mymiddleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		_, err := c.Cookie("user")
		if err != nil {
			return c.Redirect(http.StatusSeeOther, "/login")
		}
		return next(c)
	}
}
