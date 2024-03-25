package handlers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/database"
	"berjcode/dependency/dtos"
	"berjcode/dependency/helpers"
	"berjcode/dependency/mapping"
	"berjcode/dependency/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func GetAllUsers(c echo.Context) error {

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	var users []models.User
	db.Find(&users)

	return c.JSON(http.StatusOK, users)
}

func CreateUser(c echo.Context) error {

	var newUserCreateDto dtos.UserCreateDto

	if err := c.Bind(&newUserCreateDto); err != nil {
		return err
	}

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	var existingUser models.User
	if err := db.Where("user_name = ?", newUserCreateDto.UserName).Or("email = ?", newUserCreateDto.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": constant.UserExisting})
	}

	salt, err := helpers.GenerateSalt()
	if err != nil {
		return err
	}

	hashedPassword := helpers.HashPassword(newUserCreateDto.PasswordHash, salt)
	var newUser = mapping.MappingUserCreateDtoToUser(newUserCreateDto, salt, hashedPassword)

	if err := db.Create(&newUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, true)
}

type CookieData struct {
	UserName string `json:"UserName"`
}

func getUserByID(id uint) (models.User, error) {
	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return models.User{}, err
	}
	defer db.Close()

	var user models.User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return models.User{}, echo.NewHTTPError(http.StatusNotFound, constant.UserNameNotFound)
		}
		return models.User{}, err
	}

	return user, nil
}

func UpdateUser(c echo.Context) error {
	fmt.Println(c)
	var newUserUpdateDto dtos.UserUpdateDto
	if err := c.Bind(&newUserUpdateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}

	dbUser, err := getUserByID(newUserUpdateDto.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, constant.UserNameNotFound)
	}

	newPassword := helpers.HashPassword(newUserUpdateDto.PasswordHash, dbUser.Salt)
	dbUser.NameSurname = newUserUpdateDto.NameSurname
	dbUser.Email = newUserUpdateDto.Email
	dbUser.PasswordHash = newPassword

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Save(&dbUser).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": constant.UserUpdatedInfo})
}

func GetUserData(c echo.Context) error {

	paramId := c.Param("id")
	fmt.Println("id", paramId)
	convertedID, err := strconv.ParseUint(paramId, 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidLessonID)
	}

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	user, err := getUserByID(uint(convertedID))

	if err != nil {
		return err
	}

	var userDto = mapping.MappingUserToUserDto(user)

	return c.JSON(http.StatusOK, userDto)
}
