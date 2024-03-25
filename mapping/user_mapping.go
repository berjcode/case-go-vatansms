package mapping

import (
	"berjcode/dependency/common"
	"berjcode/dependency/dtos"
	"berjcode/dependency/models"
)

func MappingUserCreateDtoToUser(userCreateDto dtos.UserCreateDto, salt string, passwordHash string) models.User {

	user := models.User{
		UserName:     userCreateDto.UserName,
		NameSurname:  userCreateDto.NameSurname,
		Email:        userCreateDto.Email,
		Salt:         salt,
		PasswordHash: passwordHash,
		EntityBase: common.EntityBase{
			CreatedBy: userCreateDto.CreatedBy,
		},
	}
	return user
}

func MappingUserToUserDto(user models.User) dtos.UserDto {
	userDto := dtos.UserDto{
		ID:          user.ID,
		UserName:    user.UserName,
		NameSurname: user.NameSurname,
		Email:       user.Email,
	}

	return userDto
}
