package dtos

import (
	"time"
)

type PlanCreateDto struct {
	LessonID     uint
	StartTime    time.Time
	EndTime      time.Time
	PlanStatusID uint
	CreatedBy    string
}
