package handlers

import (
	"berjcode/dependency/database"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	password := c.FormValue("password")
	username := c.FormValue("userNameAndEmail")

	var dbUser models.User
	if err := db.Where("user_name = ? OR email = ?", username, username).First(&dbUser).Error; err != nil {
		return echo.ErrBadRequest
	}

	if !helpers.CheckPassword(password, dbUser.Salt, dbUser.PasswordHash) {
		return echo.ErrUnauthorized
	}

	cookie := helpers.GetCookie("username", dbUser.UserName, time.Now().Add(24*time.Hour))
	http.SetCookie(c.Response(), cookie)

	return c.Redirect(http.StatusSeeOther, "/plan")
}
