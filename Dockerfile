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
ARG ENV_FILE_PATH=./docker.env
ENV $(cat $ENV_FILE_PATH | xargs)

# Command to run the Go application with command-line argument
ENTRYPOINT ["./main"]
