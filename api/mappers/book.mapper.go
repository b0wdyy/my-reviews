package mappers

import (
	"api/dto"
	"api/models"
)

func BookToDto(book models.Book) dto.BookDto {
	return dto.BookDto{
		ID:         book.ID,
		Title:      book.Title,
		Pages:      book.Pages,
		Notes:      book.Notes,
		Finished:   book.Finished,
		CoverImage: book.CoverImage,
		Author:     AuthorToDto(book.Author),
		Categories: CategoriesToDto(book.Categories),
	}
}

func BooksToDto(books []models.Book) []dto.BookDto {
	var booksDto []dto.BookDto

	for _, book := range books {
		booksDto = append(booksDto, BookToDto(book))
	}

	return booksDto
}
