// handlers/student_handler.go

package handlers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/database"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {

	db, err := database.NewDB("dbconfig.json")
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

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var existingUser models.User

	if err := db.Where("user_name = ?", newUser.UserName).Or("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": constant.UserExisting})
	}

	salt, err := helpers.GenerateSalt()
	if err != nil {
		return err
	}

	hashedPassword := helpers.HashPassword(newUser.PasswordHash, salt)

	newUser.Salt = salt
	newUser.PasswordHash = hashedPassword

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newUser)
}

type CookieData struct {
	UserName string `json:"UserName"`
}

func GetUserByUsername(c echo.Context) (models.User, error) {
	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	cookie, err := c.Cookie("username")
	if err != nil {
		return models.User{}, echo.NewHTTPError(http.StatusUnauthorized, constant.UserNameNotFound)
	}
	username := cookie.Value
	var user models.User
	if err := db.Where("user_name = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return models.User{}, echo.NewHTTPError(http.StatusNotFound, constant.UserNameNotFound)
		}
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(c echo.Context) error {

	username := c.FormValue("userName")
	nameSurname := c.FormValue("nameSurname")
	email := c.FormValue("email")
	passwordHash := c.FormValue("passwordHash")
	fmt.Println("formpasswordHash: ", passwordHash)
	dbUser, err := GetUserByUsername(c)
	if err != nil {

		return echo.NewHTTPError(http.StatusNotFound, constant.UserNameNotFound)
	}

	newPassword := helpers.HashPassword(passwordHash, dbUser.Salt)
	dbUser.UserName = username
	dbUser.NameSurname = nameSurname
	dbUser.Email = email
	dbUser.PasswordHash = newPassword

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Save(&dbUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": constant.UserUpdatedInfo})
}

func GetUserIDByUserName(c echo.Context) error {
	username := c.QueryParam("username")

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var userIDs []uint
	if err := db.Model(&models.User{}).Where("user_name = ?", username).Pluck("id", &userIDs).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusOK, map[string]string{"message": constant.UserNameNotFoundForUserName})
		}
		return err
	}

	if len(userIDs) == 0 {
		return c.JSON(http.StatusOK, map[string][]uint{"userIDs": nil})
	}

	return c.JSON(http.StatusOK, map[string][]uint{"userIDs": userIDs})
}
