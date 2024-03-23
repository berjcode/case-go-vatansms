package dtos

import (
	"time"
)

type PlanUpdateDto struct {
	ID           uint
	LessonID     uint
	StartTime    time.Time
	EndTime      time.Time
	PlanStatusID uint
	UpdatedOn    *time.Time
	UpdatedBy    string
}
