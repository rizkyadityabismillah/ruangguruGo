package main


import (
    "fmt"
    "net/http"

    "github.com/rest/client"
    "github.com/rest/server"
)
func main(){
	mux := http.NewServeMux()

	//server
	mux.HandleFunc("/api/student", server.Students)
	//client
	mux.HandleFunc("/api/client", client.Students)

	fmt.Println("starting web server at localhost:8080")
	http.ListenAndServe(":8080", mux)
}