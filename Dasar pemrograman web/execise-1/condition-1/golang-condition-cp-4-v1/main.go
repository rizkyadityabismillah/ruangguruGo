package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
//   -	Jika tanggal pembelian _ticket_ tersebut adalah tanggal ganjil
//   - Jika _ticket_ yang dibeli kurang dari 5 _ticket_, maka mendapat diskon 15% dari total pembelian
//   - Jika _ticket_ yang dibeli minimal 5 _ticket_ keatas, maka mendapat diskon 25% dari total pembelian
//   - Jika tanggal pembelian _ticket_ tersebut adalah tanggal genap
//   - Jika _ticket_ yang dibeli kurang dari 5 _ticket_, maka mendapat diskon 10% dari total pembelian
//   - Jika _ticket_ yang dibeli minimal 5 _ticket_ keatas, maka mendapat diskon 20% dari total pembelian
var totalTiket = VIP+regular+student
var totalPrice = (VIP * 30)+ (regular * 20) + (student*10)

if totalPrice >= 100 {
	if day % 2 != 0{
		if totalTiket < 5 {
			return float32(totalPrice) - float32(totalPrice)* 15/100
		}else{
			return float32(totalPrice) - float32(totalPrice)*25/100
		}
	}else{
		if totalTiket < 5 {
			return float32(totalPrice) - float32(totalPrice) * 10/100
		}else{
			return float32(totalPrice) - float32(totalPrice) * 20/100
		}
	}
}
	return float32(totalPrice)
}
// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(1, 1, 1, 20))
}
