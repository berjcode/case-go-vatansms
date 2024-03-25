package mapping

import (
	"berjcode/dependency/common"
	"berjcode/dependency/dtos"
	"berjcode/dependency/models"
)

func MappingPlanCreateDtoToPlan(planCreateDto dtos.PlanCreateDto) models.Plan {

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
func MappingPlanToPlanDto(plan models.Plan) dtos.PlanDto {
	planDto := dtos.PlanDto{
		ID:           plan.ID,
		LessonID:     plan.LessonID,
		StartTime:    plan.StartTime,
		EndTime:      plan.EndTime,
		PlanStatusID: plan.PlanStatusID,
	}

	return planDto
}
