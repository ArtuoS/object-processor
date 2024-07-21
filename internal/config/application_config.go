package config

import (
	"fmt"
	"os"

	"github.com/ArtuoS/object-processor/internal/constants"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type ApplicationConfig struct {
	S3Client  *s3.S3
	SQSClient *sqs.SQS
}

func NewApplicationConfig() (*ApplicationConfig, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv(constants.AWSRegion))},
	)

	sess.Config.WithCredentials(credentials.NewStaticCredentials(os.Getenv(constants.AWSId), os.Getenv(constants.AWSSecret), ""))

	if err != nil {
		fmt.Printf("Failed to create session: %v\n", err)
		return nil, err
	}

	return &ApplicationConfig{
		S3Client:  s3.New(sess),
		SQSClient: sqs.New(sess),
	}, nil
}
