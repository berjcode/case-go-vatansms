package handlers

import (
	"berjcode/dependency/database"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	username := c.FormValue("userNameAndEmail")
	password := c.FormValue("password")

	var dbUser models.User
	if err := db.Where("user_name = ? OR email = ?", username, username).First(&dbUser).Error; err != nil {
		return echo.ErrBadRequest
	}

	fmt.Println("dbUser:", dbUser.NameSurname)

	if !helpers.CheckPassword(password, dbUser.Salt, dbUser.PasswordHash) {
		return echo.ErrUnauthorized
	}

	user := models.User{
		UserName:    dbUser.UserName,
		NameSurname: dbUser.NameSurname,
	}

	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	cookie := http.Cookie{
		Name:     "user",
		Value:    string(userData),
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}

	http.SetCookie(c.Response(), &cookie)

	return c.Redirect(http.StatusSeeOther, "/plan")
}
