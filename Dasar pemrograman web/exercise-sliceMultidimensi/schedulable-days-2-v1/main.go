package main

import "fmt"

func SchedulableDays(villager [][]int) []int {
	availableData := make(map[int]int)
	for _,v := range villager{
		for _, tanggal := range v{
			availableData[tanggal]= availableData[tanggal] + 1 
		}
	} 
	result := []int{}
	for k,v := range availableData{
		if v == len(villager){
			result = append(result, k)
		}
	}
	return result // TODO: replace this
}

func main(){
	fmt.Println(SchedulableDays([][]int{
		{7, 12, 19, 22}, 
		{12, 19, 21, 23}, 
		{7, 12, 19}, 
		{12, 19},
	}))
}
