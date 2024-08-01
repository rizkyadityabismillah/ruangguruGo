package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	// "os"
	"strings"
)

// AIModelConnector is the struct to hold the HTTP client
type AIModelConnector struct {
	Client *http.Client
}

// Inputs defines the structure of the payload sent to the AI model
type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

// Response defines the structure of the response from the AI model
type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

// CsvToSlice converts CSV data into a map[string][]string
func CsvToSlice(data string) (map[string][]string, error) {
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	if len(records) < 1 {
		return nil, errors.New("CSV data is empty")
	}

	headers := records[0]
	result := make(map[string][]string)

	for _, header := range headers {
		result[header] = []string{}
	}

	for _, record := range records[1:] {
		for i, value := range record {
			result[headers[i]] = append(result[headers[i]], value)
		}
	}

	return result, nil
}

// ConnectAIModel sends a request to the AI model and retrieves the response
func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Response{}, errors.New("failed to connect to AI model")
	}

	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func main() {
	client := &http.Client{}
	connector := AIModelConnector{Client: client}

	// Read CSV file
	data, err := ioutil.ReadFile("data-series.csv")
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Convert CSV to slice
	table, err := CsvToSlice(string(data))
	if err != nil {
		fmt.Println("Error converting CSV to slice:", err)
		return
	}

	// Get user input for query
	fmt.Println("Enter your query about the energy consumption data:")
	var query string
	fmt.Scanln(&query)

	// Create payload
	payload := Inputs{
		Table: table,
		Query: query,
	}

	// Connect to AI model
	token := "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"
	response, err := connector.ConnectAIModel(payload, token)
	if err != nil {
		fmt.Println("Error connecting to AI model:", err)
		return
	}

	// Print the response
	fmt.Println("Response from AI model:")
	fmt.Println("Answer:", response.Answer)
	fmt.Println("Coordinates:", response.Coordinates)
	fmt.Println("Cells:", response.Cells)
	fmt.Println("Aggregator:", response.Aggregator)
}
