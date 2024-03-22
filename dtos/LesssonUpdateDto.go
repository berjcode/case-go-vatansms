package dtos

import "time"

type LessonUpdateDto struct {
	ID                uint
	LessonName        string
	LessonDescription string
	UpdatedOn         *time.Time
	UpdatedBy         string
}
