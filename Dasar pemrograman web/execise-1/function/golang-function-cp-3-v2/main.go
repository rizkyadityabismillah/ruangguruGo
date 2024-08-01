package main

import (
	"fmt"
)
func WhichShortName(name1, name2 string) string{
	if len(name1) < len(name2){
		return name1
	} else if len(name1) > len(name2){
		return name2
	}else{
		if name1 < name2{
			return name1
		}else{
			return name2
		}
	}
}
func FindShortestName(names string) string {
	curName := ""
	minName := "kjkjkajljaljldlkalklakslkslaklkskkalklaklkla"
	for _, c := range names {
		if string(c) == ";" || string(c) == " " || string(c) == ","{
			minName = WhichShortName(curName, minName)
			curName = ""
		}else{
			curName += string(c)
		}
	}
		if curName != ""{
			minName = WhichShortName(curName, minName)
		}
		return minName // TODO: replace this
	}


// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
