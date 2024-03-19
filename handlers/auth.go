package handlers

import (
	"berjcode/dependency/database"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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
	if err := db.Where("user_name = ? OR email = ?", "123", "123").First(&dbUser).Error; err != nil {
		return echo.ErrBadRequest
	}

	fmt.Println("dbUser:", dbUser.NameSurname)

	if !helpers.CheckPassword(password, dbUser.Salt, dbUser.PasswordHash) {

		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["nameandsurname"] = dbUser.NameSurname
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
