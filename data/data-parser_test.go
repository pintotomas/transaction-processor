package data

import (
	"reflect"
	"testing"
	"time"
	"transaction-processor/model"
)

type testCase struct {
	name          string
	inputData     [][]string
	expectedValid bool
	expectedError bool
	expectedTx    []*model.Transaction
}

func TestDataParser(t *testing.T) {
	testCases := []testCase{
		{
			name: "ValidData",
			inputData: [][]string{
				{"id", "date", "transaction"},
				{"1", "1/15", "+23.4"},
				{"2", "1/15", "+10.1"},
				{"3", "2/15", "-5.1"},
			},
			expectedValid: true,
			expectedError: false,
			expectedTx: []*model.Transaction{
				{ID: 1, Date: time.Date(0, time.January, 15, 0, 0, 0, 0, time.UTC), Transaction: 23.4},
				{ID: 2, Date: time.Date(0, time.January, 15, 0, 0, 0, 0, time.UTC), Transaction: 10.1},
				{ID: 3, Date: time.Date(0, time.February, 15, 0, 0, 0, 0, time.UTC), Transaction: -5.1},
			},
		},
		{
			name: "InvalidHeader",
			inputData: [][]string{
				{"invalid", "header", "format"},
				{"1", "1/15", "+23.4"},
			},
			expectedValid: false,
			expectedError: true,
		},
		{
			name: "InvalidDateFormat",
			inputData: [][]string{
				{"id", "invalid_date_format", "transaction"},
				{"1", "invalid_date", "+23.4"},
			},
			expectedValid: false,
			expectedError: true,
		},
		{
			name: "InvalidTransactionFormat",
			inputData: [][]string{
				{"id", "date", "invalid_transaction_format"},
				{"1", "1/15", "invalid_transaction"},
			},
			expectedValid: false,
			expectedError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			parser := &DataParser{data: tc.inputData}
			valid := parser.Validate()
			// Test data validation
			if valid != tc.expectedValid {
				t.Errorf("Validation failed for test case %s. Expected valid: %v, got: %v", tc.name, tc.expectedValid, valid)
			}

			// Test data parsing
			if valid {
				txs, err := parser.ParseData()
				if (err != nil) != tc.expectedError {
					t.Errorf("Error mismatch for test case %s. Expected error: %v, got: %v", tc.name, tc.expectedError, err != nil)
				}

				if !reflect.DeepEqual(txs, tc.expectedTx) {
					t.Errorf("Transaction mismatch for test case %s. Expected: %v, got: %v", tc.name, tc.expectedTx, txs)
				}
			}
		})
	}
}
