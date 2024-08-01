package main

import "fmt"

// func math(n1, n2 int) int{
// 	sum := n1 + n2 * 10
// 	min := sum - 10

// 	return min
// }
// func main(){
// 	result := math(5,3)
// 	fmt.Println("result is", result)
// }

// func sayHello(name string) string{
// 	if name == ""{
// 		return "hello kamu"
// 	} else {
// 		return "hello " + name
// 	}
// }

// func main(){
// 	result := sayHello("")
// 	fmt.Println(result)

// 	fmt.Println(sayHello("john"))
// }

//MULTIPLE RETURN

// func name() (string, string){
// 	return "afrida", "maharani"
// }

// func main(){
// 	firstname, _ :=  name()
// 	fmt.Println(firstname)
// }

//NAMED RETURN MEMBERI NAMA VARIABEL PADA RETURN
// func realName()(firstname, fullname, lastname string){
// 	firstname = "afrida"
// 	fullname = "afrida maharani"
// 	lastname = "maharani"

// 	return
// }
// func main(){
// 	firstname, fullname, lastname := realName()
// 	fmt.Println(firstname, fullname, lastname)
// }

//FUNCTION AS PARAMETER

// type Filter func(string) string
// func filterName(name string, filter Filter){
// 	namedFiltered := filter(name)
// 	fmt.Println("hello", namedFiltered)
// }

// func spamFilter(name string) string{
// 	if name == "anjing"{
// 		return "kata kasar"
// 	}else{
// 		return name
// 	}
// }

// func main(){
// 	filterName("anjing",spamFilter)
// 	filterName("",spamFilter)
// }

//ANONYMOUS FUNCTION

type Blacklist func(string) bool

func regisUser(username string, blacklist Blacklist){
	if blacklist(username){
		fmt.Println("user blocked", username)
	} else{
		fmt.Println("welcome to our server", username)
	}
}

func main(){
	blacklist := func(username string) bool{
		return username == "admin"
	}
	regisUser("admin",blacklist)
	regisUser("Afrida rizky",blacklist)
}