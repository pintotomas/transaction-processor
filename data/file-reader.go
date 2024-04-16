package data

import (
	"encoding/csv"
	"os"
)

// LocalFileHandler reads data from a local CSV file
type LocalFileHandler struct {
	path string
}

// ReadData reads data from a local CSV file
func (l *LocalFileHandler) ReadData(source string) (*DataParser, error) {
	file, err := os.Open(l.path + source)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return &DataParser{data: records}, nil
}
