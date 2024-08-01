package main

import (
	"net/http"
	"time"
	"fmt"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
	hari := ""

	switch now.Weekday(){
	case time.Sunday:
		hari = "Sunday"
	case time.Monday:
		hari = "Monday"
	case time.Tuesday:
		hari = "Tuesday"
	case time.Wednesday:
		hari = "Wednesday"
	case time.Thursday:
		hari = "Thursday"
	case time.Friday:
		hari = "Friday"
	case time.Saturday:
		hari = "Saturday"
	}
	day := now.Day()
	month := now.Month()
	year := now.Year()

	result := fmt.Sprintf("%s, %d %s %d", hari, day, month, year)
	w.Write([]byte(result))
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == ""{
			w.Write([]byte("Hello there"))
		}else{
			w.Write([]byte("Hello, "+ name + "!"))
		}
		}) // TODO: replace this
}

func main() {
	// TODO: answer here
	http.ListenAndServe("localhost:8080", nil)
}
