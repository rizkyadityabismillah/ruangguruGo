package main

import "fmt"

// func main(){
// 	var a = 10
// 	result := (&a)
// 	hasil  := *result
// 	fmt.Println(result)
// 	fmt.Println(hasil)
// }

//PASSING BY VALUE DUPLIKASI DARI VARIABEL
// func main(){
// 	var nomor = 10
// 	addByOne(nomor)
// 	fmt.Println("The value of main is ", nomor)
// }
// func addByOne(number int){
// 	number += 1
// 	fmt.Println(number)
// }
func addByOne(number *int){
	*number +=1
	fmt.Println("Added by one in function:", *number)
}
func main(){
	var number int = 10  
	addByOne(&number)
	fmt.Println("The Value of Main is : ", number)
}