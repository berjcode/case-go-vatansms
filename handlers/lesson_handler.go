package handlers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/database"
	"berjcode/dependency/models"
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

func GetLessonsByUser(c echo.Context) error {
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
