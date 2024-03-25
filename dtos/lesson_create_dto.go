package dtos

type LessonCreateDto struct {
	LessonName        string
	LessonDescription string
	UserID            uint
	CreatedBy         string
}
