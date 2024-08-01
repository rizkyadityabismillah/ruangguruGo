package main

import (
  "fmt"
  "time"
  "math/rand"
)


// Write a program that starts a goroutine to print "Hello, Goroutines!" and then prints "Main function" in the main function


// Implement two separate goroutines that calculate the square and cube of a number, respectively. Your main function should wait for both goroutines to complete their calculations. Use a simple delay to ensure the main function waits.


// Implement a goroutine that attempts to read from a channel. Use select to implement a timeout, ensuring that if the channel does not receive any data within a specified time, the program prints a timeout message and exits.


 // creating a program that simulates a message passing system between a sender and multiple receivers using goroutines and channels.


func square(n int) {
    fmt.Println("Square:", n*n)
}

func cube(n int) {
    fmt.Println("Cube:", n*n*n)
}


func sender(message chan string) {

  for i := 1; i <= 5; i++ {
    data := fmt.Sprintf("message number %d", i)
    // mengirm data melalui channel
    message <- data
    time.Sleep(500 * time.Millisecond) // Simulate some work
  }
  // tutup channel karena tidak ada data lagi yang kita harus sampaikan
  close(message)

}

func receiver(id int, messages chan string, done chan bool) {
  for msg := range messages {
    fmt.Printf("Receiver %d received: %s\n", id, msg)
  }
  fmt.Printf("Receiver %d done receiving\n", id)
  done <- true
}

// func main() {

//   messages := make(chan string)
//   done := make(chan bool)

//   go sender(messages)

//   for i := 1; i < 4; i++ {
//     go receiver(i, messages, done)
//   }

//   receiversDone := 0
//   for {
//     select {
//     case <-done:
//       receiversDone++
//       if receiversDone == 3 {
//         fmt.Println("All messages received. Exiting receivers.")
//         return
//       }
//     }
//   }

// }


func main() {

	// go func() {
	// 	fmt.Println("Hello,Goroutines!")
	// }()
  
  
  // go square(4)
  // go cube(3)
  
  // fmt.Println("Main function")
  
	// time.Sleep(100 * time.Millisecond)

  datach:= make(chan int)
  timeout:= time.After(5 * time.Second) //<-chan time.time

  go func(data chan int) {
      time.Sleep(3 * time.Second)
      data <- rand.Intn(20)
  }(datach)

  select {
    case a := <- datach :
      fmt.Println("Data : ", a)
    case <- timeout:
      fmt.Println("Time Out!")    
  }
}