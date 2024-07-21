package main

import (
	"log"
	"os"
	"strconv"

	"github.com/ArtuoS/object-processor/internal/constants"
	"github.com/ArtuoS/object-processor/internal/tasks"
)

func main() {
	deleteFilesAfterUploadStr := os.Getenv(constants.DeleteFilesAfterUpload)
	deleteFilesAfterUpload, err := strconv.ParseBool(deleteFilesAfterUploadStr)
	if err != nil {
		log.Fatalf("Error parsing DELETE_FILES_AFTER_UPLOAD: %v", err)
	}

	publishSQSEventsStr := os.Getenv(constants.PublishSQSEvents)
	publishSQSEvents, err := strconv.ParseBool(publishSQSEventsStr)
	if err != nil {
		log.Fatalf("Error parsing PUBLISH_SQS_EVENTS: %v", err)
	}

	watcher := tasks.NewWatcher(os.Getenv(constants.FolderPath), deleteFilesAfterUpload, publishSQSEvents)
	watcher.Watch()
}
