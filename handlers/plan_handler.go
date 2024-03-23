package handlers

import (
	"berjcode/dependency/common"
	"berjcode/dependency/constant"
	"berjcode/dependency/database"
	"berjcode/dependency/dtos"
	"berjcode/dependency/helpers"
	"berjcode/dependency/models"
	"net/http"
	"strconv"
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

func UpdatePlan(c echo.Context) error {

	var planUpdateDto dtos.PlanUpdateDto
	if err := c.Bind(&planUpdateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}

	if err := helpers.UpdateValidatePlan(planUpdateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var existingPlan models.Plan
	if err := db.Where("id = ?", planUpdateDto.ID).First(&existingPlan).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Plan not found")
	}

	if err != nil {
		return err
	}

	var newPlan = mappingPlanUpdateDtoToPlan(planUpdateDto)

	if err := db.Save(&newPlan).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, true)
}

// func getPlanDetailById(id uint) (models.Plan, error) {

// 	db, err := database.NewDB("dbconfig.json")
// 	if err != nil {
// 		return models.Plan{}, err
// 	}
// 	defer db.Close()

// 	var planData models.Plan
// 	if err := db.First(&planData, id).Error; err != nil {
// 		return models.Plan{}, err
// 	}
// 	return planData, nil
// }

func mappingPlanUpdateDtoToPlan(planUpdateDto dtos.PlanUpdateDto) models.Plan {
	plan := models.Plan{
		LessonID:     planUpdateDto.LessonID,
		StartTime:    planUpdateDto.StartTime,
		EndTime:      planUpdateDto.EndTime,
		PlanStatusID: planUpdateDto.PlanStatusID,
		EntityBase: common.EntityBase{
			UpdatedOn: planUpdateDto.UpdatedOn,
			UpdatedBy: planUpdateDto.UpdatedBy,
		},
	}

	return plan
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

func GetPlanById(c echo.Context) error {
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

	plan, err := getPlanDetailById(uint(convertedID))

	if err != nil {
		return err
	}

	var planDto = mappingPlanToPlanDto(plan)

	return c.JSON(http.StatusOK, planDto)
}

func getPlanDetailById(id uint) (models.Plan, error) {

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return models.Plan{}, err
	}
	defer db.Close()

	var plan models.Plan
	if err := db.First(&plan, id).Error; err != nil {
		return models.Plan{}, err
	}
	return plan, nil
}

func mappingPlanToPlanDto(plan models.Plan) dtos.PlanDto {
	planDto := dtos.PlanDto{
		ID:           plan.ID,
		LessonID:     plan.LessonID,
		StartTime:    plan.StartTime,
		EndTime:      plan.EndTime,
		PlanStatusID: plan.PlanStatusID,
	}

	return planDto
}
