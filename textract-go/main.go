package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/textract"
)

func main() {
	// Initialize AWS session using your credentials
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1"), // AWS region
		// Add any other necessary config options
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
	}

	// Create a Textract client
	svc := textract.New(sess)

	// Open the PDF file
	file, err := os.Open("./xyz.pdf") // File path to your PDF document
	if err != nil {
		log.Fatalf("Failed to open PDF file: %v", err)
	}
	defer file.Close()

	// Read the file content
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Prepare input for Textract
	input := &textract.AnalyzeDocumentInput{
		Document: &textract.Document{
			Bytes: buffer,
		},
		// Specify minimal feature types as required by AWS Textract
		FeatureTypes: []*string{
			aws.String("TABLES"),
			aws.String("FORMS"),
			aws.String("SIGNATURES"),
		},
	}

	// Call Textract API to analyze the document
	result, err := svc.AnalyzeDocument(input)
	if err != nil {
		log.Fatalf("Failed to analyze document: %v", err)
	}

	// Prepare JSON formatted output
	var jsonOutput []string
	for _, page := range result.Blocks {
		if *page.BlockType == "LINE" {
			jsonOutput = append(jsonOutput, *page.Text)
		}
	}

	// Marshal JSON formatted output
	jsonBytes, err := json.MarshalIndent(jsonOutput, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal result to JSON: %v", err)
	}

	// Write JSON to a file
	outputFile, err := os.Create("./output.json")
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer outputFile.Close()

	_, err = outputFile.Write(jsonBytes)
	if err != nil {
		log.Fatalf("Failed to write JSON to file: %v", err)
	}

	fmt.Println("JSON data written to output.json")
}
