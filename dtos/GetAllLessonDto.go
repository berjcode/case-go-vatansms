package dtos

import "time"

type GetAllLessonDto struct {
	ID                uint
	LessonName        string
	LessonDescription string
	UserID            uint
	CreatedOn         *time.Time
}
