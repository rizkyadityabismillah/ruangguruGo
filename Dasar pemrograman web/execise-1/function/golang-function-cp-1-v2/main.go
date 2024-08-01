package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {
	var countMonth string
	if month == 1 {
		countMonth = "January"
	} else if month == 2{
		countMonth = "February"
	} else if month == 3{
		countMonth = "March"
	} else if month == 4{
		countMonth = "April"
	} else if month == 5 {
		countMonth = "May"
	}  else if month == 6 {
		countMonth = "June"
	} else if month == 7 {
		countMonth = "July"
	} else if month == 8 {
		countMonth = "August"
	} else if month == 9 {
		countMonth = "September"
	} else if month == 10 {
		countMonth = "October"
	} else if month == 11 {
		countMonth = "November"
	} else if month ==12 {
		countMonth = "December"
	} else{
		countMonth = "Invalid Month!"
	}
	 
	var countDay string

	if day < 10{
		countDay = "0" + fmt.Sprint(day)
	} else{
		countDay = fmt.Sprint(day)
	}
	return (countDay + "-" + countMonth + "-" + fmt.Sprint(year))
	}
	


// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(11, 11, 2012))
}
