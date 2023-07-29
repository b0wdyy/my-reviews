package services

import (
	"api/dto"
	"api/initializers"
	"api/mappers"
	"api/models"
	"api/utils"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(ctx *fiber.Ctx) []dto.BookDto {
	var books []models.Book
	dbQuery := initializers.DB.Model(&models.Book{})

	if query := ctx.Query("unfinished"); query != "" {
		dbQuery = initializers.DB.Where("finished = ?", query)
	}

	dbQuery.Order("created_at desc").Preload("Categories").Preload("Author").Find(&books)

	return mappers.BooksToDto(books)
}

func GetBook(id string) dto.BookDto {
	var book models.Book

	initializers.DB.First(&book, id).Preload("Categories").Preload("Author")

	return mappers.BookToDto(book)
}

func NewBook(c *fiber.Ctx) (dto.BookDto, error) {
	title := c.FormValue("title")
	categories := c.FormValue("categories")
	pages, _ := strconv.Atoi(c.FormValue("pages"))
	notes := c.FormValue("notes")
	finished := c.FormValue("finished") == "true"
	authorId, _ := strconv.Atoi(c.FormValue("author_id"))
	coverImage, err := c.FormFile("cover_image")

	if err != nil {
		return dto.BookDto{}, errors.New(err.Error())
	}

	image, imageError := utils.HandleImageUpload(coverImage, "books")

	if imageError != "" {
		return dto.BookDto{}, errors.New(imageError)
	}

	categoriesSlice := strings.Split(categories, ",")
	convertedAuthorId := uint(authorId)

	book := models.Book{
		Title:      title,
		Pages:      int64(pages),
		Notes:      notes,
		Finished:   finished,
		CoverImage: image,
		AuthorID:   &convertedAuthorId,
	}

	initializers.DB.Create(&book)

	var categoryModels []models.Category

	for _, categoryStr := range categoriesSlice {
		trimmedCategoryStr := strings.TrimSpace(categoryStr)
		category := models.Category{
			Description: trimmedCategoryStr,
		}

		initializers.DB.Where("lower(description) = ?", strings.ToLower(trimmedCategoryStr)).FirstOrCreate(&category)
		categoryModels = append(categoryModels, category)
	}

	initializers.DB.Model(&book).Association("Categories").Append(&categoryModels)
	initializers.DB.Save(&book)

	return mappers.BookToDto(book), nil
}

func UpdateBook(c *fiber.Ctx) (dto.BookDto, error) {
	var book models.Book

	initializers.DB.First(&book, c.Params("id"))

	title := c.FormValue("title")
	pages, _ := strconv.Atoi(c.FormValue("pages"))
	notes := c.FormValue("notes")
	finished := c.FormValue("finished") == "true"
	cover_image, err := c.FormFile("cover_image")

	if err != nil {
		return dto.BookDto{}, errors.New(err.Error())
	}

	initializers.DB.Model(&book).Updates(models.Book{
		Title:      title,
		Pages:      int64(pages),
		Notes:      notes,
		Finished:   finished,
		CoverImage: cover_image.Filename,
	})

	return mappers.BookToDto(book), nil
}

func DeleteBook(id string) {
	var book models.Book

	initializers.DB.Delete(&book, id)
}
