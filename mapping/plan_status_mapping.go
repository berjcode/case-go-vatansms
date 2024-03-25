package mapping

import (
	"berjcode/dependency/common"
	"berjcode/dependency/dtos"
	"berjcode/dependency/models"
)

func MappingPlanStatusToPlanStatusDto(planStatus models.PlanStatus) dtos.PlanStatusDto {
	planStatusDto := dtos.PlanStatusDto{
		ID:   planStatus.ID,
		Name: planStatus.Name,
	}

	return planStatusDto
}

func MappingPlanStatusToGetAllPlanStatusDto(planStatuses []models.PlanStatus) []dtos.GetAllPlanStatusDto {
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

func MappingPlanStatusCreateDtoToPlanStatus(planStatusCreateDto dtos.PlanStatusCreateDto) models.PlanStatus {

	planStatus := models.PlanStatus{
		Name: planStatusCreateDto.Name,
		EntityBase: common.EntityBase{
			CreatedBy: planStatusCreateDto.CreatedBy,
		},
	}
	return planStatus
}
