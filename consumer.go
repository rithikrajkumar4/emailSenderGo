package main

import (
	"fmt"
	"log"
	"net/smtp"
	"sync"
	"time"
)

func emailWorker(id int, ch chan Recipient, wg *sync.WaitGroup) {

	defer wg.Done()

	for recipient := range ch {
		// Simulate sending email
		smtpHost := "localhost"
		smtpPort := "1025"

		// formattedMessage := fmt.Sprintf("To: %s\r\nSubject: Test Email\r\n\r\nHello %s,\r\nThis is a test email.\r\n", recipient.Email, recipient.Name)
		// msg := []byte(formattedMessage)

		msg, msg_err := executeTemplate(recipient)

		if msg_err != nil {
			fmt.Println("Error in Executing Template", msg_err)
			continue
		}
		fmt.Printf("Worker %d: Sending email to %s\n", id, recipient.Email)

		err := smtp.SendMail(
			smtpHost+":"+smtpPort,
			nil,
			"test@yopmail.com",
			[]string{recipient.Email},
			[]byte(msg))
		if err != nil {
			log.Fatal(err)
		}

		time.Sleep(100 * time.Millisecond) // Simulate time taken to send an email

		fmt.Printf("Worker %d: Sent email to %s\n", id, recipient.Email)
	}

}
