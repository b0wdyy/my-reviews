package mappers

import (
	"api/dto"
	"api/models"
)

func CategoryToDto(category models.Category) dto.CategoryDto {
	return dto.CategoryDto{
		ID:          category.ID,
		Description: category.Description,
	}
}

func CategoriesToDto(categories []models.Category) []dto.CategoryDto {
	var categoriesDto []dto.CategoryDto

	for _, category := range categories {
		categoriesDto = append(categoriesDto, CategoryToDto(category))
	}

	return categoriesDto
}
