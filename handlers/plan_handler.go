package handlers

import (
	"berjcode/dependency/common"
	"berjcode/dependency/constant"
	"berjcode/dependency/database"
	"berjcode/dependency/dtos"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func CreatePlan(c echo.Context) error {

	var planCreateDto dtos.PlanCreateDto
	if err := c.Bind(&planCreateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}

	if err := helpers.CreateValidatePlan(planCreateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	startTime := planCreateDto.StartTime
	endTime := planCreateDto.EndTime
	ExistsCheckPlan(startTime, endTime)
	if err := ExistsCheckPlan(startTime, endTime); err != nil {
		return err
	}
	var plan = mappingPlanCreateDtoToPlan(planCreateDto)

	if err := db.Create(&plan).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, true)

}
func ExistsCheckPlan(startTime, endTime time.Time) error {

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()
	var count int64
	db.Where("start_time < ?", endTime).Where("end_time > ?", startTime).Table("plans").Count(&count)

	if err != nil {
		return err
	}
	if count > 0 {
		return echo.NewHTTPError(http.StatusBadRequest, "Belirtilen zaman aralığında zaten bir plan mevcut")
	}

	return nil
}

func mappingPlanCreateDtoToPlan(planCreateDto dtos.PlanCreateDto) models.Plan {

	lesson := models.Plan{
		LessonID:     planCreateDto.LessonID,
		PlanStatusID: planCreateDto.PlanStatusID,
		StartTime:    planCreateDto.StartTime,
		EndTime:      planCreateDto.EndTime,
		EntityBase: common.EntityBase{
			CreatedBy: planCreateDto.CreatedBy,
		},
	}
	return lesson
}
