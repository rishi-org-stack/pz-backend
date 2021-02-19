package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Set up authentication information.
	auth := smtp.PlainAuth("", "rishistack17@gmail.com", "rishijha1709", "mail.example.com")
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"sabitajha1709@example.net"}
	msg := []byte("To: recipient@example.net\r\n" +
		"Subject: discount Gophers!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	err := smtp.SendMail("ristack17@gmail.com:25", auth, "rishistack17@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
