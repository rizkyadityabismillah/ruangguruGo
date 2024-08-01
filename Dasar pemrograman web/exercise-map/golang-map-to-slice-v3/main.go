package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	result := [][]string{}
	for k, v := range mapData{
		fmt.Println("k",k)
		fmt.Println("v",v)
		element := []string{k,v}

		result = append(result, element)
	}
	return result // TODO: replace this
}
func main(){
	mapData := map[string]string{
		"hello": "world", 
		"John": "Doe", 
		"age":"14",
	}
	sliceData := MapToSlice(mapData)
	fmt.Println(sliceData)
}
