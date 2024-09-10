package mailing

import (
	"fmt"
	"net/smtp"
	"transaction-processor/model"
)

// SMTPClient mailing client
type SMTPClient struct {
	Host string
	Port string
}

// NewSMTPClient returns a new SMTPClient
func NewSMTPClient(host, port string) *SMTPClient {
	return &SMTPClient{
		Host: host,
		Port: port,
	}
}

// Send authenticates and sends email
func (s *SMTPClient) Send(email *model.Email) error {
	// authenticate
	auth := smtp.PlainAuth("", email.From, email.Credentials, s.Host)

	// send email
	err := smtp.SendMail(s.Host+":"+s.Port, auth, email.From, []string{email.To}, []byte(email.Message))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}