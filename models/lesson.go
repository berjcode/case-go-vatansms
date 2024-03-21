package models

import (
	"berjcode/dependency/common"
)

type Lesson struct {
	ID                uint   `gorm:"primary_key"`
	LessonName        string `gorm:"not null;type:nvarchar(60)"`
	LessonDescription string `gorm:"type:nvarchar(200)"`
	UserID            uint   `gorm:"not null"`
	User              User   `gorm:"foreignkey:UserID"`
	common.EntityBase
}
