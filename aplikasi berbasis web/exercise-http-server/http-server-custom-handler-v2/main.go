package main

import (
	"math/rand"
	"net/http"
)


var Quotes = []string{
	"Be yourself; everyone else is already taken. ― Oscar Wilde",
	"Be the change that you wish to see in the world. ― Mahatma Gandhi",
	"I have not failed. I've just found 10,000 ways that won't work. ― Thomas A. Edison",
	"It is never too late to be what you might have been. ― George Eliot",
	"Everything you can imagine is real. ― Pablo Picasso",
	"Nothing is impossible, the word itself says 'I'm possible'! ― Audrey Hepburn",
}
type QuotesHandler struct{
}
// TODO: answer here

func (qh QuotesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index := rand.Intn(len(Quotes))
	w.Write([]byte(Quotes[index]))
	// TODO: answer here

}

func main() {
	handler := QuotesHandler{}
	go http.ListenAndServe("localhost:8080", handler)
}
