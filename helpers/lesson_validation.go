package helpers

import (
	"berjcode/dependency/constant"
	"berjcode/dependency/dtos"
	"errors"
)

func UpdateValidateLesson(lessonDto dtos.LessonUpdateDto) error {

	if lessonDto.ID == 0 || lessonDto.LessonName == "" || lessonDto.LessonDescription == "" {
		return errors.New(constant.RequriedField)
	}

	if len(lessonDto.LessonName) < 2 || len(lessonDto.LessonName) > 100 {
		return errors.New(constant.MustBetweenTwoAndOneHundredCharacters)
	}

	if len(lessonDto.LessonDescription) < 2 || len(lessonDto.LessonDescription) > 30 {
		return errors.New(constant.MustBetweenTwoAndthirtyCharacters)
	}

	return nil
}
