package mappers

import (
	"api/dto"
	"api/models"
)

func UserToDto(user models.User) dto.UserDto {
	return dto.UserDto{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}
