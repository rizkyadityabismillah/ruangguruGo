package main

import (
	"net/http"
)

const userName = "afrida"
const passWord = "200705"

func Login(w http.ResponseWriter, r *http.Request){
	username,password,ok :=r.BasicAuth()

	if !ok{
		w.Write([]byte("something went wrong"))
		return
	}

	isValid := (username == userName) && (password == passWord)
	if !isValid {
		w.Write([]byte("Invalid Username or Password\n"))
	}
	w.Write([]byte("You are logged in!\n"))
}
func main(){
	http.HandleFunc("/login", Login)
	server := new(http.Server)
	server.Addr = ":8080"
	server.ListenAndServe()
}