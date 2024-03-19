// handlers/student_handler.go

package handlers

import (
	"net/http"
	"berjcode/dependency/database"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"

	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {

	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var users []models.User
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {
	var newUser models.User

	if err := c.Bind(&newUser); err != nil {
		return err
	}

	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	salt, err := helpers.GenerateSalt()
	if err != nil {
		return err
	}

	// Şifreyi hash'le
	hashedPassword := helpers.HashPassword(newUser.PasswordHash, salt)

	newUser.Salt = salt
	newUser.PasswordHash = hashedPassword

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newUser)
}
