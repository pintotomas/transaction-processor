package main

import (
	"fmt"
	"log"
	"transaction-processor/data"
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
	// Print the transactions
	for _, tx := range transactions {
		fmt.Printf("ID: %d, Date: %s, Transaction: %.2f\n", tx.ID, tx.Date.Format("2006-01-02"), tx.Transaction)
	}
	//
	//from := "tomasstoritest@gmail.com"
	//password := "pntu ntch dehp frtj"
	//
	//// Receiver email address.
	//to := []string{
	//	"test@gmail.com",
	//}
	//
	//// smtp server configuration.
	//smtpHost := "smtp.gmail.com"
	//smtpPort := "587"
	//
	//// Message.
	//message := []byte("This is a test email message.")
	//
	//// Authentication.
	//auth := smtp.PlainAuth("", from, password, smtpHost)
	//
	//// Sending email.
	//err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Email Sent Successfully!")

}
