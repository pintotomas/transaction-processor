package main

import (
	"log"
	"transaction-processor/data"
	"transaction-processor/mailing"
	"transaction-processor/model"
)

func main() {

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

	client := mailing.NewSMTPClient("smtp.gmail.com", "587")

	htmlContent, err := mailing.HTMLFormat(summary)

	if err != nil {
		log.Fatalf("Error formatting summary: %s", err)
		return
	}

	err = client.Send(&model.Email{
		Subject:     "Transactions Summary",
		From:        "tomasstoritest@gmail.com",
		Credentials: "pntu ntch dehp frtj",
		To:          "tomasp834@gmail.com",
		Message:     htmlContent,
	})
	if err != nil {
		log.Fatalf("Error sending email: %s", err)
		return
	}
	log.Println("Successfully sent email!")
}
