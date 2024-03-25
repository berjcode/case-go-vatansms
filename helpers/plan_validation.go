package helpers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/dtos"
	"errors"
)

func CreateValidatePlan(planCreateDto dtos.PlanCreateDto) error {

	if planCreateDto.LessonID == 0 || planCreateDto.StartTime.IsZero() || planCreateDto.EndTime.IsZero() || planCreateDto.CreatedBy == "" || planCreateDto.PlanStatusID == 0 {
		return errors.New(constant.RequriedField)
	}

	if len(planCreateDto.CreatedBy) < 2 || len(planCreateDto.CreatedBy) > 50 {
		return errors.New(constant.MustBetweenTwoAndTwentyCharacters)
	}

	if planCreateDto.EndTime.Before(planCreateDto.StartTime) {
		return errors.New("end time , startime'dan küçük olamaz")
	}

	return nil
}

func UpdateValidatePlan(planUpdateDto dtos.PlanUpdateDto) error {

	if planUpdateDto.LessonID == 0 || planUpdateDto.StartTime.IsZero() || planUpdateDto.EndTime.IsZero() || planUpdateDto.UpdatedBy == "" || planUpdateDto.PlanStatusID == 0 || planUpdateDto.UpdatedOn.IsZero() || planUpdateDto.ID == 0 {
		return errors.New(constant.RequriedField)
	}

	if len(planUpdateDto.UpdatedBy) < 2 || len(planUpdateDto.UpdatedBy) > 50 {
		return errors.New(constant.MustBetweenTwoAndTwentyCharacters)
	}

	if planUpdateDto.EndTime.Before(planUpdateDto.StartTime) {
		return errors.New("end time , startime'dan küçük olamaz")
	}

	return nil
}
