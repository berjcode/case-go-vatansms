package models

import (
	"berjcode/dependency/common"
	"time"
)

type Plan struct {
	ID              uint       `gorm:"primary_key"`
	UserID          uint       `gorm:"not null"`
	User            User       `gorm:"foreignkey:UserID"`
	PlanName        string     `gorm:"not null;type:nvarchar(100)"`
	PlanDescription string     `gorm:"type:nvarchar(200)"`
	StartTime       time.Time  `gorm:"not null"`
	EndTime         time.Time  `gorm:"not null"`
	PlanStatusID    uint       `gorm:"not null"`
	PlanStatus      PlanStatus `gorm:"foreignkey:PlanStatusID"`
	common.EntityBase
}
