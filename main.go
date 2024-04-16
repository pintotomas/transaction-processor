package main

import (
	"fmt"
	"log"
	"os"
	"transaction-processor/data"
	"transaction-processor/mailing"
	"transaction-processor/model"
)

func main() {

	// Read environment variables
	senderEmail := os.Getenv("SENDER_EMAIL")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	fmt.Println(senderEmail)
	fmt.Println(emailPassword)
	fmt.Println(smtpHost)
	fmt.Println(smtpPort)

	// Read command-line arguments
	args := os.Args[1:] // Skip the program name
	var recipientEmail string
	if len(args) == 1 {
		recipientEmail = args[0]
	} else {
		log.Fatal("Please run with <recipient_email> argument")
		return
	}

	dr, err := data.NewDataReader("local")
	if err != nil {
		log.Fatalf("Error creating data reader: %s", err)
		return
	}
	err = dr.ReadData("csv/transactions1.csv")
	if err != nil {
		log.Fatalf("Error reading data: %s", err)
		return
	}
	transactions, err := dr.ParseData()
	if err != nil {
		log.Fatalf("Error parsing data: %s", err)
		return
	}

	summary := model.CalculateSummary(transactions)

	client := mailing.NewSMTPClient(smtpHost, smtpPort)

	htmlContent, err := mailing.HTMLFormat(summary)

	if err != nil {
		log.Fatalf("Error formatting summary: %s", err)
		return
	}

	err = client.Send(&model.Email{
		Subject:     "Transactions Summary",
		From:        senderEmail,
		Credentials: emailPassword,
		To:          recipientEmail,
		Message:     htmlContent,
	})
	if err != nil {
		log.Fatalf("Error sending email: %s", err)
		return
	}
	log.Println("Successfully sent email!")
}
