package main

import "fmt"
import "strings"

//paramter variadic
func myClub(club string, player ...string){
	var myClubAsString = strings.Join(player,",")
	fmt.Println("the club is", club)
	fmt.Println("my player is",myClubAsString)
}
func main(){
	myClub("Real Madrid", "Toni Kroos",  "Bale","Ronaldo")
	myClub("Manchester City", "Bernardo", "Sergio Agüero")
}
// func tambah(){
// 	n1 := 10;
// 	n2 := 30;

// 	var jumlah = n1 + n2
// 	fmt.Println("10+30 =", jumlah)
// }
// func main(){
// 	tambah()
// }

// func sayHello(name string){
// 	fmt.Println("hello", name)
// }
// func main(){
// 	sayHello("afrida")
// 	sayHello("maharani")
// 	sayHello("afrida")
// }

// func addNumber(n1, n2 int){
// 	sum := n1 + n2
// 	fmt.Println("hasil sum", sum)
// }

// func main(){
// 	addNumber(6,10)
// }

