package models

import (
	"berjcode/dependency/common"

	"gorm.io/gorm"
)

type PlanStatus struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null;type:nvarchar(60)"`
	common.EntityBase
	gorm.Model `gorm:"TableName:plan_statuses"`
}
