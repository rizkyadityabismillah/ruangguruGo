package main

import (
	"a21hc3NpZ25tZW50/internal"
	"fmt"
	"strconv"
	"strings"
)

func AdvanceCalculator(calculate string) float32 {
	tokens := strings.Split(calculate, " ") // Membagi string calculate menjadi slice tokens menggunakan spasi sebagai pemisah

	if len(tokens) == 0 { // Memeriksa apakah slice tokens kosong
		return 0 // Mengembalikan 0 jika slice tokens kosong
	}

	number, err := strconv.ParseFloat(tokens[0], 32) // Mengkonversi string pertama dalam slice tokens menjadi float32
	if err != nil {                                  // Memeriksa jika terjadi kesalahan saat konversi
		return 0 // Mengembalikan 0 jika terjadi kesalahan saat konversi
	}

	calc := internal.NewCalculator(float32(number)) // Membuat instance dari Calculator dari package internal dengan menggunakan angka pertama sebagai nilai awal
	for i := 1; i < len(tokens); i += 2 {           // Melakukan iterasi melalui slice tokens dimulai dari indeks kedua dengan selisih 2
		operator := tokens[i]                               // Mengambil operator dari slice tokens
		operand, err := strconv.ParseFloat(tokens[i+1], 32) // Mengkonversi operand (angka setelah operator) menjadi float32
		if err != nil {                                     // Memeriksa jika terjadi kesalahan saat konversi
			return 0 // Mengembalikan 0 jika terjadi kesalahan saat konversi
		}

		switch operator { // Melakukan switch case berdasarkan operator
		case "+": // Jika operator adalah penambahan
			calc.Add(float32(operand)) // Menambahkan operand ke nilai hasil kalkulator
		case "-": // Jika operator adalah pengurangan
			calc.Subtract(float32(operand)) // Mengurangkan operand dari nilai hasil kalkulator
		case "*": // Jika operator adalah perkalian
			calc.Multiply(float32(operand)) // Mengalikan operand dengan nilai hasil kalkulator
		case "/": // Jika operator adalah pembagian
			calc.Divide(float32(operand)) // Membagi nilai hasil kalkulator dengan operand
		default: // Jika operator tidak dikenali
			return 0 // Mengembalikan 0
		}
	}

	return calc.Result() // Mengembalikan nilai hasil kalkulasi
}

func main() {
	res := AdvanceCalculator("3 * 4 / 2 + 10 - 5")

	fmt.Println(res)
}
