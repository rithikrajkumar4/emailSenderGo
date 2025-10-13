package main

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"
	"time"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {
	var startTime = time.Now()
	fmt.Println("Email Dispatcher Running...", startTime)
	recipientChannel := make(chan Recipient) // Using Unbuffered Channel
	go func() {
		err := loadRecipients("./emails.csv", recipientChannel) // Producer sending data to channel
		if err != nil {
			fmt.Println("Error in Loading Recipients", err)
		}
	}()

	workerCount := 5
	var wg sync.WaitGroup
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg) // Starting multiple consumers
	}

	wg.Wait() // Let the program run for a while to process emails
	var endTime = time.Now()
	fmt.Println("Time Taken:", endTime.Sub(startTime))
	fmt.Println("Email Dispatcher Stopped.")
}

type emailData struct {
	Name         string
	SupportEmail string
	CompanyName  string
}

func executeTemplate(recipient Recipient) (string, error) {
	// Placeholder for template execution logic
	t, err := template.ParseFiles("email.tmpl")
	if err != nil {
		return "", err
	}

	var tpl bytes.Buffer
	err = t.Execute(&tpl, emailData{
		Name:         recipient.Name,
		SupportEmail: "test@yopmail.com",
		CompanyName:  "Your Company",
	})
	if err != nil {
		return "", err
	}

	return tpl.String(), nil
}
