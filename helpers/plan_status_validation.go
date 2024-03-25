package helpers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/dtos"
	"errors"
)

func UpdateValidatePlanStatus(planStatusUpdateDto dtos.PlanStatusUpdateDto) error {

	if planStatusUpdateDto.ID == 0 || planStatusUpdateDto.Name == "" {
		return errors.New(constant.RequriedField)
	}

	if len(planStatusUpdateDto.Name) < 2 || len(planStatusUpdateDto.Name) > 50 {
		return errors.New(constant.MustBetweenTwoAndFiftyCharacters)
	}
	return nil
}

func CreateValidatePlanStatus(planStatusCreateDto dtos.PlanStatusCreateDto) error {

	if planStatusCreateDto.Name == "" || planStatusCreateDto.CreatedBy == "" {
		return errors.New(constant.RequriedField)
	}

	if len(planStatusCreateDto.Name) < 2 || len(planStatusCreateDto.Name) > 50 {
		return errors.New(constant.MustBetweenTwoAndFiftyCharacters)
	}
	return nil
}
