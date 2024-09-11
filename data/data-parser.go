package data

import (
	"strconv"
	"time"
	"transaction-processor/model"
)

// DataParser struct
type DataParser struct {
	data [][]string
}

// Validate validates the data header
func (d *DataParser) Validate() bool {
	if len(d.data) == 0 {
		return false
	}
	// fixed size array
	var expectedHeader = [...]string{idHeader, dateHeader, transactionHeader}
	// Check if header matches the expected header
	header := d.data[0]
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
func (d *DataParser) ParseData() ([]*model.Transaction, error) {
	// Create a slice to store transactions
	var transactions []*model.Transaction

	// Parse each line and create Transaction structs
	for _, line := range d.data[1:] {
		idStr := line[0]
		dateStr := line[1]
		transactionStr := line[2]

		// Parse date string
		date, err := time.Parse("1/2", dateStr)
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
