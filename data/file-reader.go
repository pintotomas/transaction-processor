package data

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
	"transaction-processor/model"
)

// LocalFileReader reads data from a local CSV file
type LocalFileReader struct {
	records [][]string
}

// ReadData reads data from a local CSV file
func (l *LocalFileReader) ReadData(source string) error {
	file, err := os.Open(source)
	defer file.Close()
	if err != nil {
		return err
	}
	// Read the CSV file
	reader := csv.NewReader(file)
	l.records, err = reader.ReadAll()
	if err != nil {
		return err
	}
	return nil
}

// Validate validates the data header
func (l *LocalFileReader) Validate() bool {
	// fixed size array
	var expectedHeader = [...]string{idHeader, dateHeader, transactionHeader}
	// Check if header matches the expected header
	header := l.records[0]
	if len(header) != len(expectedHeader) {
		return false
	}
	for i := range header {
		if header[i] != expectedHeader[i] {
			return false
		}
	}
	return true
}

// ParseData parses csv and returns transactions
func (l *LocalFileReader) ParseData() ([]*model.Transaction, error) {
	// Create a slice to store transactions
	var transactions []*model.Transaction

	// Parse each line and create Transaction structs
	for _, line := range l.records[1:] {
		idStr := line[0]
		dateStr := line[1]
		transactionStr := line[2]

		// Parse date string
		date, err := time.Parse("1/2/06", dateStr)
		if err != nil {
			return nil, err
		}

		// Parse transaction string
		transaction, err := strconv.ParseFloat(transactionStr, 64)
		if err != nil {
			return nil, err
		}

		// Parse id string to integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return nil, err
		}

		// Create Transaction struct
		tx := &model.Transaction{
			ID:          id,
			Date:        date,
			Transaction: transaction,
		}

		// Append transaction to slice
		transactions = append(transactions, tx)
	}
	return transactions, nil
}
