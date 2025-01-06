package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

const apiURL = "https://api-inference.huggingface.co/models/facebook/bart-large-cnn"

type RequestPayload struct {
	Inputs string `json:"inputs"`
}

type ResponsePayload struct {
	SummaryText string `json:"summary_text"`
}

func main() {
	// Read the Hugging Face API token from the environment variable
	apiToken := os.Getenv("HUGGING_FACE_API_TOKEN")

	// Define a flag for the input text
	textPtr := flag.String("text", "", "Text to summarize")
	flag.Parse()

	// Check if the text flag is provided
	if *textPtr == "" {
		fmt.Println("Please provide the text to summarize using the -text flag.")
		os.Exit(1)
	}

	// Input text to summarize
	text := *textPtr

	// Create the request payload
	payload := RequestPayload{Inputs: text}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling payload:", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	req.Header.Set("Authorization", "Bearer "+apiToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Parse the response
	var responsePayload []ResponsePayload
	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	// Print the summary
	if len(responsePayload) > 0 {
		fmt.Println("Summary:", responsePayload[0].SummaryText)
	} else {
		fmt.Println("No summary returned")
	}
}
