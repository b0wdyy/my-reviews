package services

import (
	"api/dto"
	"api/initializers"
	"api/mappers"
	"api/models"
	"api/utils"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GetAuthors() []dto.AuthorDto {
	var authors []models.Author
	initializers.DB.Find(&authors).Preload("Books")

	return mappers.AuthorsToDto(authors)
}

func CreateAuthor(c *fiber.Ctx) (dto.AuthorDto, error) {
	name := c.FormValue("name")
	picture, _ := c.FormFile("picture")

	image, imageError := utils.HandleImageUpload(picture, "authors")

	if imageError != "" {
		return dto.AuthorDto{}, errors.New(imageError)
	}

	author := models.Author{
		Name:    name,
		Picture: image,
	}

	result := initializers.DB.Create(&author)

	if result.Error != nil {
		return dto.AuthorDto{}, errors.New(result.Error.Error())
	}

	return mappers.AuthorToDto(author), nil
}
