package services

import (
	"api/initializers"
	"api/models"
	"context"
	"errors"
	"mime/multipart"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
)

func GetBooks(ctx *fiber.Ctx) []models.Book {
	var books []models.Book
	dbQuery := initializers.DB

	if query := ctx.Query("unfinished"); query != "" {
		dbQuery = initializers.DB.Where("finished = ?", query)
	}

	dbQuery.Order("created_at desc").Preload("Categories").Find(&books)

	return books
}

func GetBook(id string) models.Book {
	var book models.Book

	initializers.DB.First(&book, id)

	return book
}

func NewBook(c *fiber.Ctx) (models.Book, error) {
	title := c.FormValue("title")
	categories := c.FormValue("categories")
	pages, _ := strconv.Atoi(c.FormValue("pages"))
	notes := c.FormValue("notes")
	finished := c.FormValue("finished") == "true"
	cover_image, err := c.FormFile("cover_image")

	if err != nil {
		return models.Book{}, errors.New(err.Error())
	}

	image, imageError := handleImageUpload(cover_image)

	if imageError != "" {
		return models.Book{}, errors.New(imageError)
	}

	book := models.Book{
		Title:      title,
		Pages:      pages,
		Notes:      notes,
		Finished:   finished,
		CoverImage: image,
	}

	initializers.DB.Create(&book)

	categoriesSlice := strings.Split(categories, ",")

	for _, category := range categoriesSlice {
		category = strings.TrimSpace(category)

		var categoryModel models.Category

		initializers.DB.FirstOrCreate(&categoryModel, models.Category{Description: category})
		initializers.DB.Model(&book).Association("Categories").Append(&categoryModel)
	}

	return book, nil
}

func UpdateBook(c *fiber.Ctx) (models.Book, error) {
	var book models.Book

	initializers.DB.First(&book, c.Params("id"))

	title := c.FormValue("title")
	pages, _ := strconv.Atoi(c.FormValue("pages"))
	notes := c.FormValue("notes")
	finished := c.FormValue("finished") == "true"
	cover_image, err := c.FormFile("cover_image")

	if err != nil {
		return models.Book{}, errors.New(err.Error())
	}

	initializers.DB.Model(&book).Updates(models.Book{
		Title:      title,
		Pages:      pages,
		Notes:      notes,
		Finished:   finished,
		CoverImage: cover_image.Filename,
	})

	return book, nil
}

func DeleteBook(id string) {
	var book models.Book

	initializers.DB.Delete(&book, id)
}

func handleImageUpload(file *multipart.FileHeader) (string, string) {
	f, err := file.Open()

	if err != nil {
		return "", err.Error()
	}

	client := s3.NewFromConfig(initializers.AWSConfig)

	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("my-reviews-bowdy"),
		Key:    aws.String(file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if err != nil {
		return "", err.Error()
	}

	return result.Location, ""
}
