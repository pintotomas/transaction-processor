package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTransactionValidation(t *testing.T) {
	tests := []struct {
		name         string
		transaction  *Transaction
		expectedErrs map[string]string
	}{
		{
			name: "ValidTransaction",
			transaction: &Transaction{
				ID:          1,
				Date:        time.Now(),
				Transaction: 100.0,
			},
			expectedErrs: map[string]string{},
		},
		{
			name: "InvalidID",
			transaction: &Transaction{
				ID:          0,
				Date:        time.Now(),
				Transaction: 100.0,
			},
			expectedErrs: map[string]string{"ID": "Key: 'Transaction.ID' Error:Field validation for 'ID' failed on the 'gt' tag"},
		},
		{
			name: "InvalidDate",
			transaction: &Transaction{
				ID:          1,
				Transaction: 100.0,
			},
			expectedErrs: map[string]string{"Date": "Key: 'Transaction.Date' Error:Field validation for 'Date' failed on the 'required' tag"},
		},
		{
			name: "InvalidTransaction",
			transaction: &Transaction{
				ID:          1,
				Date:        time.Now(),
				Transaction: 0.0,
			},
			expectedErrs: map[string]string{"Transaction": "Key: 'Transaction.Transaction' Error:Field validation for 'Transaction' failed on the 'ne' tag"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.transaction.Validate()

			if len(tc.expectedErrs) == 0 {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)

				for _, expectedErr := range tc.expectedErrs {
					assert.Contains(t, err.Error(), expectedErr)
				}
			}
		})
	}
}

func TestCalculateSummary(t *testing.T) {
	tests := []struct {
		name            string
		transactions    []*Transaction
		expectedSummary *TransactionSummary
	}{
		{
			name:         "ZeroTransactions",
			transactions: []*Transaction{},
			expectedSummary: &TransactionSummary{
				TotalBalance:         0.0,
				AverageCredit:        0.0,
				AverageDebit:         0.0,
				TransactionsPerMonth: map[string]int{},
			},
		},
		{
			name: "OneCreditTransaction",
			transactions: []*Transaction{
				{
					ID:          1,
					Date:        time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 100.0,
				},
			},
			expectedSummary: &TransactionSummary{
				TotalBalance:  100.0,
				AverageCredit: 100.0,
				AverageDebit:  0.0,
				TransactionsPerMonth: map[string]int{
					"January": 1,
				},
			},
		},
		{
			name: "OneDebitTransaction",
			transactions: []*Transaction{
				{
					ID:          1,
					Date:        time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
					Transaction: -10.0,
				},
			},
			expectedSummary: &TransactionSummary{
				TotalBalance:  -10.0,
				AverageCredit: 0.0,
				AverageDebit:  -10.0,
				TransactionsPerMonth: map[string]int{
					"April": 1,
				},
			},
		},
		{
			name: "MultipleTransactions",
			transactions: []*Transaction{
				{
					ID:          1,
					Date:        time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 100.0,
				},
				{
					ID:          2,
					Date:        time.Date(2022, time.January, 10, 0, 0, 0, 0, time.UTC),
					Transaction: -50.0,
				},
				{
					ID:          3,
					Date:        time.Date(2022, time.February, 5, 0, 0, 0, 0, time.UTC),
					Transaction: 75.0,
				},
			},
			expectedSummary: &TransactionSummary{
				TotalBalance:  125.0,
				AverageDebit:  -50.0,
				AverageCredit: 87.5,
				TransactionsPerMonth: map[string]int{
					"January":  2,
					"February": 1,
				},
			},
		},
		{
			name: "OneTransactionPerMonth",
			transactions: []*Transaction{
				{
					ID:          1,
					Date:        time.Date(2022, time.January, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 1.0,
				},
				{
					ID:          2,
					Date:        time.Date(2022, time.February, 1, 0, 0, 0, 0, time.UTC),
					Transaction: -1.0,
				},
				{
					ID:          3,
					Date:        time.Date(2022, time.March, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 1.0,
				},
				{
					ID:          4,
					Date:        time.Date(2022, time.April, 1, 0, 0, 0, 0, time.UTC),
					Transaction: -1.0,
				},
				{
					ID:          5,
					Date:        time.Date(2022, time.May, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 1.0,
				},
				{
					ID:          6,
					Date:        time.Date(2022, time.June, 1, 0, 0, 0, 0, time.UTC),
					Transaction: -1.0,
				},
				{
					ID:          7,
					Date:        time.Date(2022, time.July, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 1.0,
				},
				{
					ID:          8,
					Date:        time.Date(2022, time.August, 1, 0, 0, 0, 0, time.UTC),
					Transaction: -1.0,
				},
				{
					ID:          9,
					Date:        time.Date(2022, time.September, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 1.0,
				},
				{
					ID:          10,
					Date:        time.Date(2022, time.October, 1, 0, 0, 0, 0, time.UTC),
					Transaction: -1.0,
				},
				{
					ID:          11,
					Date:        time.Date(2022, time.November, 1, 0, 0, 0, 0, time.UTC),
					Transaction: 1.0,
				},
				{
					ID:          12,
					Date:        time.Date(2022, time.December, 1, 0, 0, 0, 0, time.UTC),
					Transaction: -1.0,
				},
			},
			expectedSummary: &TransactionSummary{
				TotalBalance:  0.0,
				AverageDebit:  -1.0,
				AverageCredit: 1.0,
				TransactionsPerMonth: map[string]int{
					"January":   1,
					"February":  1,
					"March":     1,
					"April":     1,
					"May":       1,
					"June":      1,
					"July":      1,
					"August":    1,
					"September": 1,
					"October":   1,
					"November":  1,
					"December":  1,
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualSummary := CalculateSummary(tc.transactions)

			assert.Equal(t, tc.expectedSummary.TotalBalance, actualSummary.TotalBalance)
			assert.Equal(t, tc.expectedSummary.AverageCredit, actualSummary.AverageCredit)
			assert.Equal(t, tc.expectedSummary.AverageDebit, actualSummary.AverageDebit)
			assert.Equal(t, tc.expectedSummary.TransactionsPerMonth, actualSummary.TransactionsPerMonth)
		})
	}
}
