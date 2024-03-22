package dtos

type LessonCreateDto struct {
	ID                uint
	LessonName        string
	LessonDescription string
	UserID            uint
	CreatedBy         string
}
