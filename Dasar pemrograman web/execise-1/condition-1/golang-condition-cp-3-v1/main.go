package main

import "fmt"

func GetPredicate(math, scince, agama, inggris int)string{
	var result string
	var avg = (math+scince+agama+inggris)/4
	if(avg == 100){
		result = "sempurna"
	} else if (avg >= 90){
		result = "sangat baik"
	} else if (avg >=75){
		result = "baik"
	} else if (avg <=75){
		result = "kurang sangat"
	}
	return result
}
func main(){
	fmt.Println(GetPredicate(90,90,90,90))
}