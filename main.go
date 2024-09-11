package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
	"transaction-processor/data"
	"transaction-processor/mailing"
	"transaction-processor/model"
)

type Request struct {
	Email string `json:"email"`
}

type Response struct {
	Message string `json:"message"`
}

// ProcessAndSendEmail processes the file and sends the email
func ProcessAndSendEmail(recipient string) error {

	dr, err := data.NewDataReader("local")
	if err != nil {
		log.Fatalf("Error creating data reader: %s", err)
		return err
	}
	err = dr.ReadData("csv/transactions1.csv")
	if err != nil {
		log.Fatalf("Error reading data: %s", err)
		return err
	}
	transactions, err := dr.ParseData()
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

func HandleRequest(ctx context.Context, request Request) (Response, error) {

	err := ProcessAndSendEmail(request.Email)
	if err != nil {
		return Response{Message: "Failed to send transaction summary details"}, err
	}

	message := fmt.Sprintf("Successfully sent transaction summary details to, %s!", request.Email)
	return Response{Message: message}, nil
}

func main() {
	// Read environment variables
	env := os.Getenv("ENVIRONMENT")
	if env == "local" {
		// Read command-line arguments
		args := os.Args[1:] // Skip the program name
		var recipientEmail string
		if len(args) == 1 {
			recipientEmail = args[0]
		} else {
			log.Fatal("Please run with <recipient_email> argument")
			return
		}
		ProcessAndSendEmail(recipientEmail)
	} else if env == "production" {
		lambda.Start(HandleRequest)
	}
}
