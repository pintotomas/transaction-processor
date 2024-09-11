# Transaction Processor

## Running Instructions

### Development Environment

Developed using Go version 1.19.1.

### Using Docker

For running locally or with docker, ensure that the following environment variables are set:

```bash
ENVIRONMENT=local
SENDER_EMAIL=...
EMAIL_PASSWORD=...
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
LOCAL_FILE_PATH=csv/
```

I have provided a docker.env with those settings, but you can provide your own email and credentials, as well as choose a different SMTP host.

Build the Docker image:
```
docker build -t transaction-processor .
```

Run the Docker container:
```
docker run --env-file docker.env transaction-processor <recipient_email@example.com> <file_name.csv>
```

I have provided some files of my own for different scenarios, you can add any other file in the /csv folder 

**Note:** Once the email is sent, it's advisable to check the spam folder.

### Production Deployment

For production deployment, the application uses AWS S3 instead of local files. Follow these steps to deploy to AWS Lambda:

1. Create an ECR repository and push the Docker image there.
2. Create an S3 bucket
3. Create an AWS Lambda function from the image and ensure that the environment following variables are set:

```bash
ENVIRONMENT=production
SENDER_EMAIL=...
EMAIL_PASSWORD=...
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
AWS_S3_BUCKET=<your_S3_bucket_name>
```

Ensure that your AWS Lambda function role (in the IAM service) has access to S3 to avoid errors.

Also, make sure to upload the files to your s3 bucket

To invoke the Lambda function, test it with an event containing a request similar to the following:

```json
{
  "email": "recipient@test.com",
  "file-name": "..."
}
```