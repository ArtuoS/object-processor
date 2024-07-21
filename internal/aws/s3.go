package aws

import (
	"os"
	"path/filepath"

	"github.com/ArtuoS/object-processor/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	bucket string
	config *config.ApplicationConfig
}

func NewS3(bucket string, config *config.ApplicationConfig) *S3 {
	return &S3{
		bucket: bucket,
		config: config,
	}
}

func (s *S3) Upload(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = s.config.S3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filepath.Base(path)),
		Body:   file,
	})
	if err != nil {
		return err
	}

	return nil
}
