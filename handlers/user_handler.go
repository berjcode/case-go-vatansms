// handlers/student_handler.go

package handlers

import (
	"berjcode/dependency/database"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
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

	var existingUser models.User

	if err := db.Where("user_name = ?", newUser.UserName).Or("email = ?", newUser.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bu kullanıcı zaten mevcut"})
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
	db, err := database.NewDB()
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	cookie, err := c.Cookie("username")
	if err != nil {
		return models.User{}, echo.NewHTTPError(http.StatusUnauthorized, "Kullanıcı adı bulunamadı")
	}
	username := cookie.Value
	fmt.Println("xx", username)
	var user models.User
	if err := db.Where("user_name = ?", username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return models.User{}, echo.NewHTTPError(http.StatusNotFound, "Kullanıcı bulunamadı")
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

		return echo.NewHTTPError(http.StatusNotFound, "Kullanıcı bulunamadı")
	}

	newPassword := helpers.HashPassword(passwordHash, dbUser.Salt)
	fmt.Println("Salt: ", dbUser.Salt)
	fmt.Println("newPassword: ", newPassword)
	dbUser.UserName = username
	dbUser.NameSurname = nameSurname
	dbUser.Email = email
	dbUser.PasswordHash = newPassword

	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Save(&dbUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Kullanıcı bilgileri güncellendi"})
}
