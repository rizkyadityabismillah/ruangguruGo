package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	var result []int
	result = make([]int, 0)
	// for i := 1; i <= 31;i++ {
	// 	isAmamAvailable := false
	// 	isPermanaAvailable := false
	// 	for _, date := range date1 {
	// 		if date == i {
	// 			isAmamAvailable = true
	// 		}
	// 	}
	// 	for _, date := range date2 {
	// 		if date == i {
	// 			isPermanaAvailable = true
	// 		}
	// 	}
	// 	if isAmamAvailable && isPermanaAvailable {
	// 		result = append(result, i)
	// 	}
	// }
	for _, imamDate := range date1 {
		for _,permanaDate := range date2{
			if imamDate == permanaDate{
				result = append(result, imamDate)
			}
		}
	}
	return result // TODO: replace this
}
func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4},[]int{3,4,5}))
}
