package models

import (
	"berjcode/dependency/common"
	"time"
	"gorm.io/gorm"
)

type Plan struct {
	ID           uint       `gorm:"primary_key"`
	UserID       uint       `gorm:"not null"`
	User         User       `gorm:"foreignkey:UserID"`
	LessonID     uint       `gorm:"not null"`
	Lesson       Lesson     `gorm:"foreignkey:LessonID"`
	StartTime    time.Time  `gorm:"not null"`
	EndTime      time.Time  `gorm:"not null"`
	PlanStatusID uint       `gorm:"not null"`
	PlanStatus   PlanStatus `gorm:"foreignkey:PlanStatusID"`
	gorm.Model `gorm:"TableName:plans"`
	common.EntityBase
}
