package main

import "fmt"

func GraduateStudent(score int, absent int) string {
	var status string
	if score >= 70 && absent < 5{
		status = "lulus"
	} else if score >= 70 || absent >= 5{
		status = "tidak lulus"
	} else{
		status = "tidak lulus"
	}
	return status
 // TODO: replace this
	}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GraduateStudent(80, 5))
}
