package handlers

import (
	"berjcode/dependency/constant"
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

	var loginForm models.UserLogin
	if err := c.Bind(&loginForm); err != nil {
		return echo.ErrBadRequest
	}

	if err := helpers.ValidateLoginForm(loginForm); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var dbUser models.User
	if err := db.Where("user_name = ? OR email = ?", loginForm.UsernameAndEmail, loginForm.UsernameAndEmail).First(&dbUser).Error; err != nil {
		return echo.ErrBadRequest
	}

	if !helpers.CheckPassword(loginForm.Password, dbUser.Salt, dbUser.PasswordHash) {
		return echo.ErrUnauthorized
	}

	cookie := helpers.GetCookie("username", dbUser.UserName, time.Now().Add(24*time.Hour))
	http.SetCookie(c.Response(), cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": constant.SuccessLogin})
}
