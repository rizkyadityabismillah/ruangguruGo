// package main

// import "fmt"

// func main(){
// 	numb := make(chan int)
// 	s := []int{7,2,8-9,4,10}
// 	go sum(s[:len(s)/2], numb)
// 	go sum(s[len(s)/2:], numb)

// 	resulta := <-numb
// 	resultb := <-numb

// 	fmt.Println(resulta + resultb)
// }
// func sum(s []int, c chan<- int){
// 	res := 0
// 	for _,v := range s{
// 		res += v
// 	}
// 	c <- res
// }