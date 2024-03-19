package models

import (
	"berjcode/dependency/common"

	"gorm.io/gorm"
)

type Lesson struct {
	ID                uint   `gorm:"primary_key"`
	LessonName        string `gorm:"not null;type:nvarchar(60)"`
	LessonDescription string `gorm:"type:nvarchar(200)"`
	gorm.Model        `gorm:"TableName:lessons"`
	common.EntityBase
}
