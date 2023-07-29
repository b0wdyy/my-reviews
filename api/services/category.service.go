package services

import (
	"api/dto"
	"api/initializers"
	"api/mappers"
	"api/models"
	"errors"
	"strings"
)

func CreateCategory(description string) (dto.CategoryDto, error) {
	category := models.Category{
		Description: description,
	}
	result := initializers.DB.Where("lower(description) = ?", strings.ToLower(description)).FirstOrCreate(&category)

	if result.RowsAffected == 0 {
		return dto.CategoryDto{}, errors.New("category already exists")
	}

	return mappers.CategoryToDto(category), nil
}
