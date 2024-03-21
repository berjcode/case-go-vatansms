package models

import (
	"berjcode/dependency/common"
	"time"
)

type Plan struct {
	ID           uint       `gorm:"primary_key"`
	LessonID     uint       `gorm:"not null"`
	Lesson       Lesson     `gorm:"foreignkey:LessonID"`
	StartTime    time.Time  `gorm:"not null"`
	EndTime      time.Time  `gorm:"not null"`
	PlanStatusID uint       `gorm:"not null"`
	PlanStatus   PlanStatus `gorm:"foreignkey:PlanStatusID"`
	common.EntityBase
}
