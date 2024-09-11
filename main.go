package main

import (
	"errors"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
	"transaction-processor/data"
	"transaction-processor/mailing"
	"transaction-processor/model"
)

// ProcessAndSendEmail processes the file and sends the email
func ProcessAndSendEmail(recipient string, file string) error {

	dr, err := data.NewDataReader(os.Getenv("ENVIRONMENT"))
	if err != nil {
		log.Fatalf("Error creating data reader: %s", err)
		return err
	}
	dp, err := dr.ReadData(file)
	if err != nil {
		log.Fatalf("Error reading data: %s", err)
		return err
	}
	if !dp.Validate() {
		log.Fatal("Failed to validate file")
		return errors.New("invalid file")
	}
	transactions, err := dp.ParseData()
	if err != nil {
		log.Fatalf("Error parsing data: %s", err)
		return err
	}

	summary := model.CalculateSummary(transactions)

	client := mailing.NewSMTPClient(os.Getenv("SMTP_HOST"), os.Getenv("SMTP_PORT"))

	htmlContent, err := mailing.HTMLFormat(summary)

	if err != nil {
		log.Fatalf("Error formatting summary: %s", err)
		return err
	}

	err = client.Send(&model.Email{
		Subject:     "Transactions Summary",
		From:        os.Getenv("SENDER_EMAIL"),
		Credentials: os.Getenv("EMAIL_PASSWORD"),
		To:          recipient,
		Message:     htmlContent,
	})
	if err != nil {
		log.Fatalf("Error sending email: %s", err)
		return err
	}
	log.Println("Successfully sent email!")
	return nil
}

func main() {
	// Read environment variables
	env := os.Getenv("ENVIRONMENT")
	switch env {
	case "local":
		// Read command-line arguments
		args := os.Args[1:] // Skip the program name
		var recipientEmail string
		var fileName string
		if len(args) == 2 {
			recipientEmail = args[0]
			fileName = args[1]
		} else {
			log.Fatal("Please run with <recipient_email> <file_name> arguments")
			return
		}
		ProcessAndSendEmail(recipientEmail, fileName)
	case "production":
		lambda.Start(HandleRequest)
	default:
		panic("Unknown environment. Please set ENVIRONMENT to 'local' or 'production'")
	}
}
