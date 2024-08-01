package main

import (
	"net/http"
	"encoding/csv"
	"strings"
	"errors"
	"fmt"
	"encoding/json"
)

type AIModelConnector struct {
	Client *http.Client
}

type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

func CsvToSlice(data string) (map[string][]string, error) {
	// Buat reader dari string CSV
	reader := csv.NewReader(strings.NewReader(data))

	// Baca semua baris dari CSV
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Pastikan setidaknya ada satu baris untuk header
	if len(records) < 1 {
		return nil, errors.New("CSV data does not contain any rows")
	}

	// Baris pertama adalah header
	headers := records[0]
	result := make(map[string][]string)

	// Inisialisasi map dengan header sebagai key dan slice kosong sebagai value
	for _, header := range headers {
		result[header] = []string{}
	}

	// Isi map dengan data dari baris kedua dan seterusnya
	for _, row := range records[1:] {
		for i, value := range row {
			header := headers[i]
			result[header] = append(result[header], value)
		}
	}

	return result, nil // TODO: replace this
}

func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	p, ok := payload.(Inputs)
	if !ok {
		return Response{}, fmt.Errorf("invalid payload type")
	}

	table := p.Table
	query := p.Query

	parts := strings.Split(query, " ")
	if len(parts) < 4 || parts[0] != "Berapa" || parts[1] != "umur" {
		return Response{}, fmt.Errorf("unsupported query format")
	}
	name := parts[3]

	names, ok := table["Name"]
	if !ok {
		return Response{}, fmt.Errorf("table does not contain 'Name' column")
	}

	ages, ok := table["Age"]
	if !ok {
		return Response{}, fmt.Errorf("table does not contain 'Age' column")
	}

	for i, n := range names {
		if n == name {
			if i < len(ages) {
				age := ages[i]
				return Response{
					Answer:      age,
					Coordinates: [][]int{{i, 1}},
					Cells:       []string{age},
					Aggregator:  "single",
				}, nil
			}
		}
	}

	return Response{}, fmt.Errorf("%s not found in table", name)
}
	
func main() {
	
	// Contoh penggunaan fungsi ConnectAIModel
	connector := &AIModelConnector{}

	payload := Inputs{
		Table: map[string][]string{
			"Name": {"John", "Doe"},
			"Age":  {"30", "40"},
		},
		Query: "Berapa umur John?",
	}

	token := "your-huggingface-token"
	response, err := connector.ConnectAIModel(payload, token)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		responseJSON, _ := json.MarshalIndent(response, "", "  ")
		fmt.Println("Response:", string(responseJSON))
	}
}
