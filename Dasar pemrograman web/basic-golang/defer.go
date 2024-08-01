package main 

import "fmt"
//DEFER ADALAH FUNCTION YANG BISA DIJADWALKAN EKSEKUSINYA SETELAH SEBUAH FUNCTION SELESAI DI EKSEKUSI
// DEFER AKAN SELALU DI EKSEKUSI WALAUPUUN DISEBUAH FUNCTION ADA ERROR
func logging(){
	fmt.Println("selesai memanggil function")
}

func runApplication(value int){
	defer logging()
	fmt.Println("aplikasi berjalan")
	result := 10/value
	fmt.Println("result", result)
}

func main(){
	runApplication(0)
}