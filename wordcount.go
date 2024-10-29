package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

// Function to count words in a string
func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func main() {
	// Get the file path from command-line argument
	filePath := flag.String("file", "", "Path to the CSV file")
	flag.Parse()

	// Check if the file path is provided
	if *filePath == "" {
		log.Fatalf("Please provide the path to a CSV file using the -file flag.")
	}

	// Open the CSV file
	file, err := os.Open(*filePath)
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all data from CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	// Initialize total word count
	totalWordCount := 0

	// Process each entry in the CSV
	for i, record := range records {
		if len(record) > 0 {
			entry := record[0]
			wordCount := countWords(entry)
			fmt.Printf("Entry %d word count: %d\n", i+1, wordCount)
			totalWordCount += wordCount
		}
	}

	// Print total word count
	fmt.Printf("Total word count: %d\n", totalWordCount)
}
