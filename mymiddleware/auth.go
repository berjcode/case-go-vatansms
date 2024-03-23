package mymiddleware

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookieHeader := c.Request().Header.Get("Cookie")

		cookies := strings.Split(cookieHeader, ";")

		var found bool
		for _, cookie := range cookies {
			if strings.Contains(cookie, "userdata") {
				found = true
				break
			}
		}

		if !found {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
