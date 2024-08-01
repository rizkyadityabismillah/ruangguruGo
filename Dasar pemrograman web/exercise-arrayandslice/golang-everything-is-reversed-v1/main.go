package main

import (
	"fmt"
	"strconv"
)

func ReverseArray(arr [5]int) [5]int {
	result := [5] int{}
	for i :=0; i < len(arr); i++{
		result[i]= arr[len(arr)-1-i]
	}
	return  result 
}

func ReverseDigit(number int)int{
	s:= strconv.Itoa(number)

	rev := ""
	for _,c := range s{
		rev = string(c)+ rev
	}
	result,_ := strconv.Atoi(rev)
	return result
}
func ReverseData(arr [5]int) [5]int{
	a := ReverseArray(arr)

	for i :=0; i < len(a); i++{
		a[i]=ReverseDigit(a[i])
	}
	return a
}
func main(){
	fmt.Println(ReverseData([5]int {456789, 44332, 2221, 12, 10}))
}