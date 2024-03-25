package mapping

import (
	"berjcode/dependency/common"
	"berjcode/dependency/dtos"
	"berjcode/dependency/models"
)

func MappingLessonToLessonDto(lesson models.Lesson) dtos.LessonDto {
	lessonDto := dtos.LessonDto{
		ID:                lesson.ID,
		LessonName:        lesson.LessonName,
		LessonDescription: lesson.LessonDescription,
		UserID:            lesson.UserID,
	}

	return lessonDto
}

func MappingLessonCreateDtoToLesson(lessonCreateDto dtos.LessonCreateDto) models.Lesson {

	lesson := models.Lesson{
		LessonName:        lessonCreateDto.LessonName,
		LessonDescription: lessonCreateDto.LessonDescription,
		UserID:            lessonCreateDto.UserID,
		EntityBase: common.EntityBase{
			CreatedBy: lessonCreateDto.CreatedBy,
		},
	}
	return lesson
}

func MappingLessonToGetAllLessonDto(lessons []models.Lesson) []dtos.GetAllLessonDto {
	var getAllLessonDtos []dtos.GetAllLessonDto
	for _, lesson := range lessons {
		dto := dtos.GetAllLessonDto{
			ID:                lesson.ID,
			LessonName:        lesson.LessonName,
			LessonDescription: lesson.LessonDescription,
			UserID:            lesson.UserID,
			CreatedOn:         lesson.CreatedOn,
		}
		getAllLessonDtos = append(getAllLessonDtos, dto)
	}

	return getAllLessonDtos
}
