package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
)

// Transaction transaction struct
type Transaction struct {
	ID          int       `validate:"gt=0"`
	Date        time.Time `validate:"required"`
	Transaction float64   `validate:"ne=0"`
}

// Validate validates transaction
func (t *Transaction) Validate() error {
	validate := validator.New()
	return validate.Struct(t)
}

func CalculateSummary(transactions []*Transaction) *TransactionSummary {
	summary := &TransactionSummary{
		TransactionsPerMonth: make(map[string]int),
	}

	var totalDebit, totalCredit float64
	creditCount := 0
	debitCount := 0

	for _, tx := range transactions {
		// Update total balance
		summary.TotalBalance += tx.Transaction

		// Update total debit and credit amounts
		if tx.Transaction > 0 {
			totalCredit += tx.Transaction
			creditCount++
		} else {
			totalDebit += tx.Transaction
			debitCount++
		}

		// Count transactions per month
		month := tx.Date.Month().String()
		summary.TransactionsPerMonth[month]++

	}

	// Calculate averages
	if creditCount > 0 {
		summary.AverageCredit = totalCredit / float64(creditCount)
	}
	if debitCount > 0 {
		summary.AverageDebit = totalDebit / float64(debitCount)
	}

	return summary
}

type TransactionSummary struct {
	TotalBalance         float64
	AverageDebit         float64
	AverageCredit        float64
	TransactionsPerMonth map[string]int
}

func (ts *TransactionSummary) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("\r\nTotal balance is %.2f", ts.TotalBalance))
	builder.WriteString(fmt.Sprintf("\r\nAverage debit amount: %.2f", ts.AverageDebit))
	builder.WriteString(fmt.Sprintf("\r\nAverage credit amount: %.2f\n", ts.AverageCredit))

	// Define the order of months
	monthOrder := []string{
		"January", "February", "March", "April", "May", "June",
		"July", "August", "September", "October", "November", "December",
	}

	// Append month data in specified order
	for _, month := range monthOrder {
		if count, ok := ts.TransactionsPerMonth[month]; ok {
			builder.WriteString(fmt.Sprintf("\r\nNumber of transactions in %s: %d", month, count))
		}
	}

	return builder.String()
}
