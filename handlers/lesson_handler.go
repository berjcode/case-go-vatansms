package handlers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/database"
	"berjcode/dependency/dtos"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUserLesson(c echo.Context) error {

	var newLesson models.Lesson
	userID := c.FormValue("userID")
	lessonName := c.FormValue("lessonName")
	lessonDescription := c.FormValue("lessonDescription")

	userIDInt, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err
	}
	newLesson.UserID = uint(userIDInt)
	newLesson.LessonName = lessonName
	newLesson.LessonDescription = lessonDescription

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var existingLesson models.Lesson
	if err := db.Where("lesson_name = ?", lessonName).Where("user_id = ?", newLesson.UserID).First(&existingLesson).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": constant.ErrorMessageExistingLesson})
	}

	if err := db.Create(&newLesson).Error; err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/plan")
}

func GetAllLessonsByUser(c echo.Context) error {
	userIDStr := c.QueryParam("userID")
	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var lessons []models.Lesson
	if err := db.Where("user_id = ?", userIDStr).Find(&lessons).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, lessons)
}

func UpdateLesson(c echo.Context) error {

	var lessonDtoUpdate dtos.LessonUpdateDto
	if err := c.Bind(&lessonDtoUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}
	fmt.Println(lessonDtoUpdate)
	if err := helpers.UpdateValidateLesson(lessonDtoUpdate); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	lessonData, err := getLessonDetailById(lessonDtoUpdate.ID)

	if err != nil {
		return err
	}

	lessonData.LessonName = lessonDtoUpdate.LessonName
	lessonData.LessonDescription = lessonDtoUpdate.LessonDescription

	if err := db.Save(&lessonData).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{"message": constant.UserUpdatedInfo})
}

func GetLessonById(c echo.Context) error {
	paramId := c.QueryParam("id")
	convertedID, err := strconv.ParseUint(paramId, 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidLessonID)
	}

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	lesson, err := getLessonDetailById(uint(convertedID))

	if err != nil {
		return err
	}

	var lessonDto = mappingLessonToLessonDto(lesson)

	return c.JSON(http.StatusOK, lessonDto)
}

func mappingLessonToLessonDto(lesson models.Lesson) dtos.LessonDto {
	lessonDto := dtos.LessonDto{
		ID:                lesson.ID,
		LessonName:        lesson.LessonName,
		LessonDescription: lesson.LessonDescription,
		UserID:            lesson.UserID,
	}

	return lessonDto
}

func getLessonDetailById(id uint) (models.Lesson, error) {

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return models.Lesson{}, err
	}
	defer db.Close()

	var lessonData models.Lesson
	if err := db.First(&lessonData, id).Error; err != nil {
		return models.Lesson{}, err
	}
	return lessonData, nil
}
