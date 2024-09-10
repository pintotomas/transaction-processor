package data

import (
	"errors"
	"transaction-processor/model"
)

const (
	idHeader          = "id"
	dateHeader        = "date"
	transactionHeader = "transaction"
)

// DataReader defines the interface for reading data
type DataReader interface {
	ReadData(source string) error
	Validate() bool
	ParseData() ([]*model.Transaction, error)
}

// NewDataReader creates a new data reader based on the source type
func NewDataReader(sourceType string) (DataReader, error) {
	switch sourceType {
	case "local":
		return &LocalFileReader{}, nil
	case "production":
		return &S3FileReader{}, nil
	default:
		return nil, errors.New("unsupported data source type")
	}
}
