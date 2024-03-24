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

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	exists, err := checkPlanStatusByName(planStatusCreateDto.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, constant.ErrorDatabase)
	}
	if exists {
		return c.JSON(http.StatusOK, map[string]string{constant.Message: constant.ExistsRegisterPlanStatus})
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

	if err := existsCheckCountPlanStatus(planStatusUpdateDto.ID); err != nil {
		return err
	}

	exists, err := checkPlanStatusByName(planStatusUpdateDto.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, constant.ErrorDatabase)
	}
	if exists {
		return c.JSON(http.StatusOK, map[string]string{constant.Message: constant.ExistsRegisterPlanStatus})
	}

	var planStatus = mappingPlanStatusUpdateToPlanStatus(planStatusUpdateDto)

	if err := db.Update(&planStatus).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, true)
}

func GetPlanStatusById(c echo.Context) error {
	paramId := c.Param("id")
	convertedID, err := strconv.ParseUint(paramId, 10, 64)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidLessonID)
	}

	db, err := database.NewDB(constant.DbConfig)
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
	db, err := database.NewDB(constant.DbConfig)
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

// private

func existsCheckCountPlanStatus(id uint) error {

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	var count int64
	db.Where("id = ?", id).Table("plan_statuses").Count(&count)

	if err != nil {
		return err
	}
	if count == 0 {
		return echo.NewHTTPError(http.StatusBadRequest, constant.NotExistsRegisterPlanStatus)
	}

	return nil
}

func getPlanStatusDetailById(id uint) (models.PlanStatus, error) {

	db, err := database.NewDB(constant.DbConfig)
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

func checkPlanStatusByName(planName string) (bool, error) {

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return false, err
	}
	defer db.Close()

	var count int
	err = db.Model(&models.PlanStatus{}).Where("name = ?", planName).Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// mapping
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

func mappingPlanStatusUpdateToPlanStatus(planStatusUpdateDto dtos.PlanStatusUpdateDto) models.PlanStatus {
	planStatus := models.PlanStatus{
		Name: planStatusUpdateDto.Name,
		EntityBase: common.EntityBase{
			UpdatedOn: planStatusUpdateDto.UpdatedOn,
			UpdatedBy: planStatusUpdateDto.UpdatedBy,
		},
	}
	return planStatus
}
