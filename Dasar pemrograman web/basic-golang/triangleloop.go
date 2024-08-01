package main

import "fmt"

func main(){
	tinggi := 5;
	for (i := 1 ; i <= tinggi ; i++){
		for(j := 0; j <= i; i++){
			fmt.Print(" ")
		}
		fmt.Print("*")
	}
}