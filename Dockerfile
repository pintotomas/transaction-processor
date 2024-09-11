# Use official Golang image version 1.19.1 as the base image
FROM golang:1.19.1

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Set environment variables from docker.env file
ENV ENVIRONMENT=local
ENV SENDER_EMAIL=tomasstoritest@gmail.com
ENV EMAIL_PASSWORD="pntu ntch dehp frtj"
ENV SMTP_HOST=smtp.gmail.com
ENV SMTP_PORT=587
ENV AWS_REGION=
ENV AWS_S3_BUCKET=
ENV LOCAL_FILE_PATH=csv/
ENV AWS_ACCESS_KEY_ID=
ENV AWS_SECRET_ACCESS_KEY=

# Command to run the Go application with command-line argument
ENTRYPOINT ["./main"]
