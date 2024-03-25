package handlers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/database"
	"berjcode/dependency/dtos"
	"berjcode/dependency/helpers"
	"berjcode/dependency/mapping"
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

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	startTime := planCreateDto.StartTime
	endTime := planCreateDto.EndTime
	existsCheckPlanByTime(startTime, endTime)
	if err := existsCheckPlanByTime(startTime, endTime); err != nil {
		return err
	}
	var plan = mapping.MappingPlanCreateDtoToPlan(planCreateDto)

	if err := db.Create(&plan).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, true)

}

func UpdatePlan(c echo.Context) error {

	var planUpdateDto dtos.PlanUpdateDto
	if err := c.Bind(&planUpdateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, constant.InvalidRequestValid)
	}

	if err := helpers.UpdateValidatePlan(planUpdateDto); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()

	var existingPlan models.Plan
	if err := db.Where("id = ?", planUpdateDto.ID).First(&existingPlan).Error; err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Plan not found")
	}

	startTime := planUpdateDto.StartTime
	endTime := planUpdateDto.EndTime
	existsCheckPlanByTime(startTime, endTime)
	if err := existsCheckPlanByTime(startTime, endTime); err != nil {
		return err
	}

	existingPlan.LessonID = planUpdateDto.LessonID
	existingPlan.StartTime = planUpdateDto.StartTime
	existingPlan.EndTime = planUpdateDto.EndTime
	existingPlan.PlanStatusID = planUpdateDto.PlanStatusID
	existingPlan.UpdatedBy = planUpdateDto.UpdatedBy
	existingPlan.UpdatedOn = planUpdateDto.UpdatedOn

	if err := db.Save(&existingPlan).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusNoContent, true)
}

func GetPlanDetails(c echo.Context) error {
	userID := c.Param("userid")
	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err
	}

	var planDetails []models.PlanDetail

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Raw(`
        SELECT 
            pl.id, 
            ls.lesson_name, 
			DATE_FORMAT(pl.start_time, '%Y-%m-%d') AS start_date, 
            DAYNAME(pl.start_time) AS start_day, 
			DATE_FORMAT(pl.start_time, '%H:%i:%s') AS start_time,
            DAYNAME(pl.end_time) AS end_day, 
			DATE_FORMAT(pl.end_time, '%Y-%m-%d') AS end_date, 
			DATE_FORMAT(pl.end_time, '%H:%i:%s') AS end_time,
            ps.name AS plan_status_name, 
            pl.created_on 
        FROM 
            newdb.plans AS pl
        INNER JOIN 
            newdb.lessons AS ls ON pl.lesson_id = ls.id
        INNER JOIN 
            newdb.plan_statuses AS ps ON pl.plan_status_id = ps.id
		where ls.user_id = ?
		ORDER BY
    		pl.start_time ASC;

    `, convertedUserID).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var planDetail models.PlanDetail
		if err := rows.Scan(&planDetail.ID, &planDetail.LessonName, &planDetail.StartDay, &planDetail.StartDate, &planDetail.StartTime, &planDetail.EndDate, &planDetail.EndDay, &planDetail.EndTime, &planDetail.PlanStatusName, &planDetail.CreatedOn); err != nil {
			return err
		}
		planDetails = append(planDetails, planDetail)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, planDetails)
}

func GetPlanDetailsByWhere(c echo.Context) error {
	userID := c.Param("userid")
	startTime := c.Param("starttime")
	endTime := c.Param("endtime")

	convertedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return err
	}

	var planDetails []models.PlanDetail

	db, err := database.NewDB(constant.DbConfig)
	if err != nil {
		return err
	}
	defer db.Close()
	rows, err := db.Raw(`
        

    SELECT 
    pl.id, 
    ls.lesson_name, 
    DATE_FORMAT(pl.start_time, '%Y-%m-%d') AS start_date, 
    DAYNAME(pl.start_time) AS start_day, 
    DATE_FORMAT(pl.start_time, '%H:%i:%s') AS start_time,
    DAYNAME(pl.end_time) AS end_day, 
    DATE_FORMAT(pl.end_time, '%Y-%m-%d') AS end_date, 
    DATE_FORMAT(pl.end_time, '%H:%i:%s') AS end_time,
    ps.name AS plan_status_name, 
    pl.created_on 
FROM 
    newdb.plans AS pl
INNER JOIN 
    newdb.lessons AS ls ON pl.lesson_id = ls.id
INNER JOIN 
    newdb.plan_statuses AS ps ON pl.plan_status_id = ps.id
WHERE 
    ls.user_id = ? AND
    pl.start_time >= ? AND 
    pl.end_time <= ?
ORDER BY
    pl.start_time ASC;


    `, convertedUserID, startTime, endTime).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var planDetail models.PlanDetail
		if err := rows.Scan(&planDetail.ID, &planDetail.LessonName, &planDetail.StartDay, &planDetail.StartDate, &planDetail.StartTime, &planDetail.EndDate, &planDetail.EndDay, &planDetail.EndTime, &planDetail.PlanStatusName, &planDetail.CreatedOn); err != nil {
			return err
		}
		planDetails = append(planDetails, planDetail)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, planDetails)
}

func GetPlanById(c echo.Context) error {
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

	plan, err := getPlanDetailById(uint(convertedID))

	if err != nil {
		return err
	}

	var planDto = mapping.MappingPlanToPlanDto(plan)

	return c.JSON(http.StatusOK, planDto)
}

// private
func getPlanDetailById(id uint) (models.Plan, error) {

	db, err := database.NewDB(constant.DbConfig)
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

func existsCheckPlanByTime(startTime, endTime time.Time) error {

	db, err := database.NewDB(constant.DbConfig)
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
		return echo.NewHTTPError(http.StatusBadRequest, constant.NotAlreadyExistsWithTime)
	}

	return nil
}
