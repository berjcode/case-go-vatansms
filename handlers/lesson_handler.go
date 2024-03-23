package handlers

import (
	"berjcode/dependency/common"
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

	var newLesson dtos.LessonCreateDto
	if err := c.Bind(&newLesson); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}

	
	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var existingLesson models.Lesson
	if err := db.Where("lesson_name = ?", newLesson.LessonName).Where("user_id = ?", newLesson.UserID).First(&existingLesson).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": constant.ErrorMessageExistingLesson})
	}
	var lesson = mappingLessonCreateDtoToLesson(newLesson)
	if err := db.Create(&lesson).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, true)
}

func GetAllLessonsByUser(c echo.Context) error {
	userIDStr := c.Param("userid")
	fmt.Println(userIDStr)
	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var newLesson []dtos.GetAllLessonDto
	var lessons []models.Lesson
	if err := db.Where("user_id = ?", userIDStr).Find(&lessons).Error; err != nil {
		return err
	}

	newLesson = mappingLessonToGetAllLessonDto(lessons)

	return c.JSON(http.StatusOK, newLesson)
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
	lessonData.UpdatedBy = lessonDtoUpdate.UpdatedBy
	lessonData.UpdatedOn = lessonDtoUpdate.UpdatedOn
	if err := db.Save(&lessonData).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, true)
}

func GetLessonById(c echo.Context) error {
	paramId := c.Param("id")
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

func mappingLessonToLessonDto(lesson models.Lesson) dtos.LessonDto {
	lessonDto := dtos.LessonDto{
		ID:                lesson.ID,
		LessonName:        lesson.LessonName,
		LessonDescription: lesson.LessonDescription,
		UserID:            lesson.UserID,
	}

	return lessonDto
}

func mappingLessonCreateDtoToLesson(lessonCreateDto dtos.LessonCreateDto) models.Lesson {

	lesson := models.Lesson{
		ID:                lessonCreateDto.ID,
		LessonName:        lessonCreateDto.LessonName,
		LessonDescription: lessonCreateDto.LessonDescription,
		UserID:            lessonCreateDto.UserID,
		EntityBase: common.EntityBase{
			CreatedBy: lessonCreateDto.CreatedBy,
		},
	}
	return lesson
}

func mappingLessonToGetAllLessonDto(lessons []models.Lesson) []dtos.GetAllLessonDto {
	var getAllLessonDtos []dtos.GetAllLessonDto
	for _, lesson := range lessons {
		dto := dtos.GetAllLessonDto{
			ID:                lesson.ID,
			LessonName:        lesson.LessonName,
			LessonDescription: lesson.LessonDescription,
			UserID:            lesson.UserID,
			CreatedOn:         lesson.CreatedOn,
		}
		getAllLessonDtos = append(getAllLessonDtos, dto)
	}

	return getAllLessonDtos
}
