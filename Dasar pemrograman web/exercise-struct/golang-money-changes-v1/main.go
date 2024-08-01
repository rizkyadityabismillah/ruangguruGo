package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}
func ExchangeCoin(amount int) []int {
	result := make([]int, 0)
	for amount > 0 {
		if amount >= 1000{
			result = append(result,  1000)
			amount -= 1000
		}else if amount >= 500{
			result = append(result, 500)
			amount -=500
		}else if amount >= 200{
			result = append(result, 200)
			amount -= 200
		} else if amount >= 100{
			result = append(result, 100)
			amount-=100
		} else if amount >= 50 {
			result = append(result, 50)
			amount -= 50
		} else if amount >= 20 {
			result = append(result, 20)
			amount -=20
		}else if amount >= 10 {
			result = append(result, 10)
			amount -=10
		} else if amount >= 5 {
			result = append(result, 5)
			amount -=5
		}else{
			result = append(result, 1)
			amount -=1
		}
	}
	return result
}
func MoneyChanges(amount int, products []Product) []int {
	totalProductCoast := 0

	for _, product := range products {
		totalProductCoast += product.Price + product.Tax
	}
	change := amount - totalProductCoast
	return ExchangeCoin(change)
}
func main() {
	products := []Product{
		{Name : "Teh", Price : 2000, Tax: 0},
		{Name : "Kopi", Price : 3000, Tax: 500},
		{Name :  "Susu", Price : 10000, Tax: 0},
	}
	changes := MoneyChanges(20000, products)
	fmt.Println(changes)
}
