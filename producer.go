package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipients(filepath string, ch chan Recipient) error {

	defer close(ch) // Channel will be closed once the function completes

	file, file_err := os.Open(filepath)

	if file_err != nil {
		fmt.Println("Error maybe in Reading File")
		return file_err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, records_err := csvReader.ReadAll()
	if records_err != nil {
		fmt.Println("Error maybe in Parsing File")
		return records_err
	}
	for _, record := range records[1:] {
		// fmt.Println("Name:", record[0], "Email:", record[1])
		ch <- Recipient{Name: record[0], Email: record[1]}
	}
	return nil
}
