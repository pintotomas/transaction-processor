package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	ID          int
	Date        time.Time
	Transaction float64
}

func main() {
	// Open the CSV file
	file, err := os.Open("data/transactions1.csv")
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer file.Close()

	// Parse the CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV: %s", err)
	}

	// Create a slice to store transactions
	var transactions []*Transaction

	// Parse each line and create Transaction structs
	for i, line := range lines {
		if i == 0 {
			// Skip header line
			continue
		}
		idStr := line[0]
		dateStr := line[1]
		transactionStr := line[2]

		// Parse date string
		date, err := time.Parse("1/2/06", dateStr)
		if err != nil {
			log.Fatalf("Error parsing date: %s", err)
		}

		// Parse transaction string
		transaction, err := strconv.ParseFloat(transactionStr, 64)
		if err != nil {
			log.Fatalf("Error parsing transaction amount: %s", err)
		}

		// Parse id string to integer
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalf("Error parsing ID: %s", err)
		}

		// Create Transaction struct
		tx := &Transaction{
			ID:          id,
			Date:        date,
			Transaction: transaction,
		}

		// Append transaction to slice
		transactions = append(transactions, tx)
	}

	// Print the transactions
	for _, tx := range transactions {
		fmt.Printf("ID: %d, Date: %s, Transaction: %.2f\n", tx.ID, tx.Date.Format("2006-01-02"), tx.Transaction)
	}

	// Sender data.
	from := "tomasstoritest@gmail.com"
	password := "pntu ntch dehp frtj"

	// Receiver email address.
	to := []string{
		"test@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")

}
