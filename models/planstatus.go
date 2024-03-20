package models

import (
	"berjcode/dependency/common"
)

type PlanStatus struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null;type:nvarchar(60)"`
	common.EntityBase
}
