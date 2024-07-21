#### How it works?
Object Processor observes a folder and simply send it to a S3 Bucket, it also send a event to a SQS queue if needed.
![a](https://github.com/ArtuoS/object-processor/blob/main/static/design.png)

#### Running Object Processor
Create an .env file in the project root with the following variables:
```
FOLDER_PATH=/app/data
S3_BUCKET=your-s3-bucket
DELETE_FILES_AFTER_UPLOAD=true/false
PUBLISH_SQS_EVENTS=true/false
SQS_QUEUE=sqs-queue-url
AWS_REGION=aws-region
AWS_ID=user-aws-id
AWS_SECRET=user-aws-secret-key
DEVICE_PATH=folder-to-observe
```
