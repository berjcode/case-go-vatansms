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

	"github.com/labstack/echo/v4"
)

func CreatePlanStatus(c echo.Context) error {

	var planStatusCreateDto dtos.PlanStatusCreateDto
	if err := c.Bind(&planStatusCreateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}

	if err := helpers.CreateValidatePlanStatus(planStatusCreateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var existingPlanStatus models.PlanStatus
	if err := db.Where("name = ?", planStatusCreateDto.Name).First(&existingPlanStatus).Error; err == nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": constant.ErrorMessageExistingLesson})
	}
	var planStatus = mappingPlanStatusCreateDtoToPlanStatus(planStatusCreateDto)
	if err := db.Create(&planStatus).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, true)
}

func UpdatePlanStatus(c echo.Context) error {

	var planStatusUpdateDto dtos.PlanStatusUpdateDto
	if err := c.Bind(&planStatusUpdateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}

	if err := helpers.UpdateValidatePlanStatus(planStatusUpdateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	planStatus, err := getPlanStatusDetailById(planStatusUpdateDto.ID)

	if err != nil {
		return err
	}

	planStatus.Name = planStatusUpdateDto.Name
	planStatus.UpdatedBy = planStatusUpdateDto.UpdatedBy
	planStatus.UpdatedOn = planStatusUpdateDto.UpdatedOn

	if err := db.Save(&planStatus).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, true)
}

func getPlanStatusDetailById(id uint) (models.PlanStatus, error) {

	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return models.PlanStatus{}, err
	}
	defer db.Close()

	var planStatus models.PlanStatus
	if err := db.First(&planStatus, id).Error; err != nil {
		return models.PlanStatus{}, err
	}
	return planStatus, nil
}

func GetPlanStatusById(c echo.Context) error {
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

	planStatus, err := getPlanStatusDetailById(uint(convertedID))

	if err != nil {
		return err
	}

	var planStatusDto = mappingPlanStatusToPlanStatusDto(planStatus)

	return c.JSON(http.StatusOK, planStatusDto)
}

func GetAllPlanStatus(c echo.Context) error {
	db, err := database.NewDB("dbconfig.json")
	if err != nil {
		return err
	}
	defer db.Close()

	var getAllPlanStatusDto []dtos.GetAllPlanStatusDto
	var planStatus []models.PlanStatus
	if err := db.Find(&planStatus).Error; err != nil {
		return err
	}

	getAllPlanStatusDto = mappingPlanStatusToGetAllPlanStatusDto(planStatus)

	return c.JSON(http.StatusOK, getAllPlanStatusDto)
}

func mappingPlanStatusToPlanStatusDto(planStatus models.PlanStatus) dtos.PlanStatusDto {
	planStatusDto := dtos.PlanStatusDto{
		ID:   planStatus.ID,
		Name: planStatus.Name,
	}

	return planStatusDto
}

func mappingPlanStatusToGetAllPlanStatusDto(planStatuses []models.PlanStatus) []dtos.GetAllPlanStatusDto {
	var getAllPlanStatusDto []dtos.GetAllPlanStatusDto
	for _, planStatus := range planStatuses {
		dto := dtos.GetAllPlanStatusDto{
			ID:   planStatus.ID,
			Name: planStatus.Name,
		}
		getAllPlanStatusDto = append(getAllPlanStatusDto, dto)
	}

	return getAllPlanStatusDto
}

func mappingPlanStatusCreateDtoToPlanStatus(planStatusCreateDto dtos.PlanStatusCreateDto) models.PlanStatus {

	planStatus := models.PlanStatus{
		Name: planStatusCreateDto.Name,
		EntityBase: common.EntityBase{
			CreatedBy: planStatusCreateDto.CreatedBy,
		},
	}
	return planStatus
}
