package main

import "fmt"

func main(){
	// for i:=1 ; i <= 10 ; i++{
	// 	fmt.Println("hello world", i)
	// } 
	// for i := 100 ; i > 0 ; i -=20{
	// 	fmt.Println(i)
	// }

	// for i , total := 1,0 ; i <=5 ; i++{
	// 	total += i
	// 	fmt.Println(total)	
	//}	
	// https://neurons.ruangguru.com/course/backend-development/dasar-pemrograman-backend/looping
		for i:=0 ; i <= 5 ; i++{
			if i % 2 == 0{
				continue
			}
			fmt.Println("hello world",i)
		}
}