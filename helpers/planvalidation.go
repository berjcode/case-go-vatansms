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

	return nil
}
