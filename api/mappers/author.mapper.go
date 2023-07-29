package mappers

import (
	"api/dto"
	"api/models"
)

func AuthorToDto(author models.Author) dto.AuthorDto {
	return dto.AuthorDto{
		ID:      author.ID,
		Name:    author.Name,
		Picture: author.Picture,
	}
}

func AuthorsToDto(authors []models.Author) []dto.AuthorDto {
	var authorsDto []dto.AuthorDto

	for _, author := range authors {
		authorsDto = append(authorsDto, AuthorToDto(author))
	}

	return authorsDto
}
