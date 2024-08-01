package main
import "fmt"

func main(){
	runtime.GOMAXPROCS(2)
	// deklarasi channel
	var message = make(chan string)

	var SayHelloTo = func (who string){
		var data = fmt.Sprintf( "Hello %s", who )
		
		message <- data // send data ke channel 'message'
	}
	// jalankan fungsi 'sayHelloTo' dalam goroutine (concurrent)
	go SayHelloTo("Afrida")
	go SayHelloTo("Rizky")
	go SayHelloTo("Maharani")

	// receive data dari channel 'messages'
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}