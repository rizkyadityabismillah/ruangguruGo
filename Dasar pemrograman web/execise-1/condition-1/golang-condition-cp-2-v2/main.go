package main

import "fmt"

func BMICalculator(gender string, height int) float64 {
	var result float64
	if gender == "laki-laki"{
		result = (float64(height-100)) - (float64(height-100)* 10/100)
	} else if gender == "perempuan"{
		result = (float64(height-100)) - (float64(height-100)* 15/100)
	}
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 170))
	fmt.Println(BMICalculator("perempuan", 165))
}
