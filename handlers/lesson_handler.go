package handlers

import (
	"berjcode/dependency/database"
	"berjcode/dependency/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateUserLesson(c echo.Context) error {

	var newLesson models.Lesson

	lessonName := c.FormValue("lessonName")
	userID := c.FormValue("userID")
	lessonDescription := c.FormValue("lessonDescription")

	userIDInt, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err
	}
	newLesson.UserID = uint(userIDInt)
	newLesson.LessonName = lessonName
	newLesson.LessonDescription = lessonDescription

	db, err := database.NewDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var existingLesson models.Lesson
	if err := db.Where("lesson_name = ?", lessonName).Where("user_id = ?", newLesson.UserID).First(&existingLesson).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Bu ders zaten mevcut"})
	}

	if err := db.Create(&newLesson).Error; err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/plan")
}
