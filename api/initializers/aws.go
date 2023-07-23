package initializers

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var AWSConfig aws.Config

func InitAWSConfig() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return
	}

	AWSConfig = cfg
}
