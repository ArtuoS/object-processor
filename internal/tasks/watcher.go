package tasks

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ArtuoS/object-processor/internal/aws"
	"github.com/ArtuoS/object-processor/internal/config"
	"github.com/ArtuoS/object-processor/internal/constants"
	"github.com/fsnotify/fsnotify"
)

type Watcher struct {
	Path                   string
	DeleteFilesAfterUpload bool
	PublishSQSEvents       bool
}

func NewWatcher(path string, deleteFilesAfterUpload, publishEvents bool) *Watcher {
	return &Watcher{
		Path:                   path,
		DeleteFilesAfterUpload: deleteFilesAfterUpload,
		PublishSQSEvents:       publishEvents,
	}
}

func (w *Watcher) Watch() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	config, err := config.NewApplicationConfig()
	if err != nil {
		log.Fatal(err)
	}

	s3 := aws.NewS3(os.Getenv(constants.S3Bucket), config)
	sqs := aws.NewSQS(os.Getenv(constants.SQSQueue), config)
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Create == fsnotify.Create {
					w.handleFileCreation(event, s3, sqs)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	err = watcher.Add(w.Path)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func (w *Watcher) handleFileCreation(event fsnotify.Event, s3 *aws.S3, sqs *aws.SQS) {
	err := s3.Upload(event.Name)
	if err != nil {
		fmt.Println(err)
	}

	if w.DeleteFilesAfterUpload {
		err = os.Remove(event.Name)
		if err != nil {
			fmt.Println(err)
		}
	}

	if w.PublishSQSEvents {
		err = sqs.Publish(&aws.SQSEvent{
			Type: "created",
			Data: "File was created.",
			Time: time.Now(),
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
