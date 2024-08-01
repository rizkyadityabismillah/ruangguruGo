package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/", getTamplate)
	http.ListenAndServe(":8080", nil)
}

func getTamplate(w http.ResponseWriter, r *http.Request){
	template := `
	<h1>Aditira</h1>
	<p>My Hobby is:</p>
	<ul>
	  <li>Programming</li>
	  <li>Gaming</li>
	</ul>
	`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(template))
}

// package main

// import (
//     "encoding/json"
//     "net/http"
// )

// type Profile struct {
//     Name    string
//     Hobbies []string
// }

// func main() {
//     http.HandleFunc("/", foo)
//     http.ListenAndServe(":8080", nil)
// }

// func foo(w http.ResponseWriter, r *http.Request) {
//     profile := Profile{"Aditira", []string{"programming", "gaming"}}

//     js, err := json.Marshal(profile)
//     if err != nil {
//         http.Error(w, err.Error(), http.StatusInternalServerError)
//         return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     w.Write(js)
// }