package data

import (
	"encoding/csv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"strings"
)

// S3FileReader reads data from an S3 bucket
type S3FileReader struct {
	awsRegion string
	bucket    string
}

// ReadData reads data from an S3 bucket
func (r *S3FileReader) ReadData(key string) (*DataParser, error) {
	// start new s3 sess
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(r.awsRegion),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	// Input parameters for GetObject API
	input := &s3.GetObjectInput{
		Bucket: aws.String(r.bucket),
		Key:    aws.String(key),
	}

	// Get the object from S3
	result, err := svc.GetObject(input)
	if err != nil {
		return nil, err
	}
	// Read the contents of the object
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	// Close the body to release the connection
	result.Body.Close()

	// Create a CSV reader from the body of the object
	reader := csv.NewReader(strings.NewReader(string(body)))

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err

	}

	// Implement logic to read data from an S3 bucket
	return &DataParser{data: records}, nil
}
