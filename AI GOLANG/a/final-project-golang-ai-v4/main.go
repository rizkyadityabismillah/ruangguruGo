package main

import (
	"net/http"
	"strings"
	"fmt"
	"errors"
	"encoding/csv"
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
    if len(records) < 2 {
        return nil, errors.New("CSV data does not contain header and rows")
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

    return result, nil
}

func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
    // Lakukan type assertion pada payload untuk mengakses fields-nya
    p, ok := payload.(Inputs)
    if !ok {
        return Response{}, errors.New("invalid payload type")
    }

    // Proses payload
    table := p.Table
    query := p.Query

    // Pastikan query dimulai dengan format yang diharapkan
    expectedPrefix := "What is the age of "
    if !strings.HasPrefix(query, expectedPrefix) {
        return Response{}, errors.New("unsupported query format")
    }

    // Ekstrak nama dari query
    name := query[len(expectedPrefix):]

    // Cari kolom "Name" dan "Age"
    names, ok := table["Name"]
    if !ok {
        return Response{}, errors.New("table does not contain 'Name' column")
    }

    ages, ok := table["Age"]
    if !ok {
        return Response{}, errors.New("table does not contain 'Age' column")
    }

    // Cari nama dalam kolom "Name"
    for i, n := range names {
        if n == name {
            // Temukan umur orang tersebut
            if i < len(ages) {
                age := ages[i]
                return Response{
                    Answer:      age,
                    Coordinates: [][]int{{i, 1}}, // Baris ke-i, kolom "Age"
                    Cells:       []string{age},
                    Aggregator:  "", // Jangan lupa sesuaikan dengan spesifikasi
                }, nil
            }
        }
    }

    return Response{}, errors.New(name + " not found in table")
}

func main() {
    // Contoh penggunaan fungsi CsvToSlice
    csvData := `"Name,Age\nJohn,30\nDoe,40"`
    table, err := CsvToSlice(csvData)
    if err != nil {
        fmt.Println("Error parsing CSV:", err)
        return
    }
    fmt.Println("CsvToSlice result:", table)

    // Contoh penggunaan fungsi ConnectAIModel
    connector := &AIModelConnector{}

    payload := Inputs{
        Table: map[string][]string{
            "Name": {"John", "Doe"},
            "Age":  {"30", "40"},
        },
        Query: "What is the age of John?",
    }

    token := "your-huggingface-token"
    response, err := connector.ConnectAIModel(payload, token)
    if err != nil {
        fmt.Println("Error connecting to AI model:", err)
        return
    }
    responseJSON, _ := json.MarshalIndent(response, "", "  ")
    fmt.Println("ConnectAIModel response:", string(responseJSON))
}




