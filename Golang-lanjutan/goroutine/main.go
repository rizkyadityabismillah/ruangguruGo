// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func greet(str string){
// // 	fmt.Println("HAI", str)
// // }
// // func main(){
// // 	fmt.Println("Main Started")
// // 	go greet("Afrida Sayang")
// // 	go greet("Sayang Afrida")
// // 	fmt.Println("Before Sleep")
// // 	time.Sleep(10 * time.Millisecond)
// // 	fmt.Println("Main ENded")
// // }

// var start time.Time

// func init(){
// 	start = time.Now()
// }
// func main(){
// 	// Sequential: Program akan menunggu tiap pemanggilan APICall sehingga banyak waktu yang diperlukan untuk menunggu
// 	ApiCallA()
// 	ApiCallB()
// 	// Concurrent: Tambahkan go saat melakukan APICall:
//     // go APICallA()
//     // go APICallB()
// 	time.Sleep(200 * time.Millisecond)
// 	fmt.Println("from main function at time", time.Since(start))
// }
// func ApiCallA(){
// 	time.Sleep(100 *  time.Millisecond)
// 	fmt.Println("API CallA Time AT", time.Since(start))
// }
// func ApiCallB(){
// 	time.Sleep(100* time.Millisecond)
// 	fmt.Println("API CALLB TIME AT", time.Since(start))
// }