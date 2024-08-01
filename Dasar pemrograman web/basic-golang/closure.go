package main

import "fmt"

// fungsi yang menggunakan variabel di luar scope atau area fungsi 
// func main(){
// 	name := "Afrida"
// 	counter := 0

// 	increment := func(){
// 		fmt.Println(name)
// 		fmt.Println("increment")
// 		counter++
// 	}
// 	increment()
// 	fmt.Println(counter)
// }

func newCounter() func()int{
	i :=0
	inc := func()int{
		i++
		return i
	}
	return inc
}

func main(){
	control := newCounter()
	fmt.Println(control())
	fmt.Println(control())
	fmt.Println(control())
	fmt.Println(control())
}