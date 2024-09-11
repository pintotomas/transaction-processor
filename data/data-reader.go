package data

import (
	"errors"
	"os"
)

const (
	idHeader          = "id"
	dateHeader        = "date"
	transactionHeader = "transaction"
)

// DataReader defines the interface for reading data
type DataReader interface {
	ReadData(source string) (*DataParser, error)
}

// NewDataReader creates a new data reader based on the source type
func NewDataReader(sourceType string) (DataReader, error) {
	switch sourceType {
	case "local":
		return &LocalFileHandler{
			path: os.Getenv("LOCAL_FILE_PATH"),
		}, nil
	case "production":
		return &S3FileReader{
			awsRegion: os.Getenv("AWS_REGION"),
			bucket:    os.Getenv("AWS_S3_BUCKET"),
		}, nil
	default:
		return nil, errors.New("unsupported data source type")
	}
}
