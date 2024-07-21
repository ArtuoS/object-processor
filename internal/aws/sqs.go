package aws

import (
	"encoding/json"
	"time"

	"github.com/ArtuoS/object-processor/internal/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSEvent struct {
	Type string
	Data string
	Time time.Time
}

type SQS struct {
	queue  string
	config *config.ApplicationConfig
}

func NewSQS(queue string, config *config.ApplicationConfig) *SQS {
	return &SQS{
		queue:  queue,
		config: config,
	}
}

func (s *SQS) Publish(event *SQSEvent) error {
	json, err := json.Marshal(event)
	if err != nil {
		return err
	}

	_, err = s.config.SQSClient.SendMessage(&sqs.SendMessageInput{
		MessageBody: aws.String(string(json)),
		QueueUrl:    aws.String(s.queue),
	})
	return err
}
