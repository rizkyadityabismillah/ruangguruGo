package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := ""
	yangPertama := true
	for _, v := range data {
		mengandungInput := strings.Contains(v, input)
		if mengandungInput {
			if yangPertama{
				result += v
				yangPertama = false
			}else{
				result += ","+ v				
			}
		}
	}
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone","iphone Bakso", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
