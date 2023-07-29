package utils

import (
	"api/initializers"
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func HandleImageUpload(file *multipart.FileHeader, folder string) (string, string) {
	f, err := file.Open()

	if err != nil {
		return "", err.Error()
	}

	client := s3.NewFromConfig(initializers.AWSConfig)

	uploader := manager.NewUploader(client)

	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("my-reviews-bowdy"),
		Key:    aws.String(folder + "/" + file.Filename),
		Body:   f,
		ACL:    "public-read",
	})

	if err != nil {
		return "", err.Error()
	}

	return result.Location, ""
}
