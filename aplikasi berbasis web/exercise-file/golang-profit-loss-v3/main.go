package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {
	content,err := os.ReadFile(path)
	if err != nil{
		return nil, err
	}
	text := string(content)
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return []string{},nil
	}
	if len(lines) == 1 && lines[0] ==""{
		return []string{},nil
	}
	return lines, nil 
}

func CalculateProfitLoss(data []string) string {
	profit := 0
	lastDate := ""

	for _, line := range data{
		tokens := strings.Split(line,";")
		date := tokens[0]
		lastDate = date
		transactionsType := tokens[1]
		amount, _ := strconv.Atoi(tokens[2])

		if transactionsType == "income"{
			profit+=amount
		}else if(transactionsType=="expense"){
			profit-=amount
		}
	}
	if profit >= 0 {
		return fmt.Sprintf("%s;profit;%d", lastDate ,profit)
	}else {
		return fmt.Sprintf("%s;loss;%d", lastDate,-profit)
	}
	
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
