package main

import (
	"fmt"
	"strconv"
	"strings"
)
func PopulationData(data []string) []map[string]any {
	result := []map[string]any{}

	for _,v := range data{
		curResult := make(map[string]any)
		fmt.Println(v)
		tokens := strings.Split(v, ";")
		Name := tokens[0]
		Age,_ := strconv.Atoi(tokens[1])
		Address := tokens[2]
		Height,_ := strconv.ParseFloat(tokens[3],64)
		IsMarried,_ := strconv.ParseBool(tokens[4])
		fmt.Println("Name", Name)
		fmt.Println("Age", Age)
		fmt.Println("Address", Address)
		fmt.Println("Height ", Height)
		fmt.Println("IsMarried ", IsMarried)
		fmt.Println()
		curResult["name"] = Name
		curResult["age"] = Age
		curResult["address"] = Address

		if tokens[3] != "" {
			curResult["height"] = Height
		}
		if tokens[4] !="" {
			curResult["isMarried"]=IsMarried
		}
		fmt.Println("CurResult", curResult)
		result = append(result, curResult)
	}
	return result // TODO: replace this
}
func main(){
	data := []string{"Budi;23;Jakarta;;", 
	"Joko;30;Bandung;;true", 
	"Susi;25;Bogor;165.42;"}
	output := PopulationData(data)
	switch val := output[1]["Age"].(type){
	case int:
		fmt.Println(val + 1)
	}
}