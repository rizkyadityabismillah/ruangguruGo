package main

import "fmt"

func TicketPlayground(height, age int) int {
// - Jika Umur di bawah 5 tahun maka tidak dapat membeli tiket wahana dengan menampilkan `-1` (repersentasi dari `error`)
// - Jika anak umur 5-7 tahun atau tinggi lebih dari 120 cm maka akan dikenakan tarif Rp 15.000.
// - Jika anak umur 8-9 tahun atau tinggi lebih dari 135 cm maka akan dikenakan tarif Rp 25.000.
// - Jika anak umur 10-11 tahun atau tinggi lebih dari 150 cm maka akan dikenakan tarif Rp 40.000.
// - Jika anak umur 12 tahun atau tinggi lebih dari 160 cm, akan dikenakan tarif Rp 60.000.
// - Jika di atas 12 tahun, akan mendapat tiket khusus Remaja yaitu sebesar Rp 100.000
if age < 5 {
		return -1
	}
	var tarif = 5000
	if age >= 5 || height >= 120{
		tarif += 10000
	} 
	if age >= 8 || height >= 135 {
		tarif += 10000
	}
	if age >= 10 || height >= 150 {
		tarif += 15000
	}
	if age >= 12 || height >= 160 {
		tarif += 20000
	}
	if age > 12{
		tarif = 100000
	}
	return tarif
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
