package dtos

import (
	"time"
)

type PlanDto struct {
	ID           uint
	LessonID     uint
	StartTime    time.Time
	EndTime      time.Time
	PlanStatusID uint
}
