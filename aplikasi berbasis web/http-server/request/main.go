// package main

// import (
//     "bytes"
//     "encoding/json"
//     "fmt"
//     "io"
//     "net/http"
// )

// // Membuat fungsi requestClient() untuk menghandle request dengan parameter Method, URL dan Body:
// func requestClient(method string, url string, body io.Reader) (string, error) {
//     // Dengan menggunakan NewRequest() kita lakukan set request yang menggunakan parameter Method, URL dan Body:
//     req, err := http.NewRequest(method, url, body)
//     if err != nil {
//         return "", err
//     }
//     // Kita juga gunakan fungsi Header untuk set tipe data apa yang di request:
//     req.Header.Set("Accept", "application/json")

//     // Kita lakukan configurasi di bawah ini untuk mengirim HTTP request ke Server:
//     client := &http.Client{}
//     resp, err := client.Do(req)
//     if err != nil {
//         return "", err
//     }

//     // Tutup response body di akhir fungsi
//     defer resp.Body.Close()

//     // Baca seluruh response body sebagai output
//     b, err := io.ReadAll(resp.Body)
//     if err != nil {
//         return "", err
//     }
// 	// Jadikan response body sebagai string dan return
//     return string(b), nil
// }

// func main() {
//     // Gunakan fungsi requestClient() yang dibuat untuk melakukan request, di sini kita gunakan method GET:
//     resp, err := requestClient("GET", "https://icanhazdadjoke.com", nil)
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println(resp)


//     // Gunakan fungsi requestClient() yang dibuat untuk melakukan request, di sini kita gunakan method GET dengan URL yang memiliki parameter yaitu `?title=bleach`:
//     resp, err = requestClient("GET", "https://animechan.vercel.app/api/quotes/anime?title=bleach", nil)
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println(resp)

//     // Set request body sebagai data yang akan di kirim ke server melalui method POST
//     postBody, _ := json.Marshal(map[string]string{
//         "name":  "aditira",
//         "email": "aditira@gmail.com",
//     })
//     responseBody := bytes.NewBuffer(postBody)

//     // Gunakan fungsi requestClient() yang dibuat untuk melakukan request, di sini kita gunakan method POST:
//     resp, err = requestClient("POST", "https://postman-echo.com/post", responseBody)
//     if err != nil {
//         panic(err)
//     }

//     fmt.Println(resp)

//     // Kita juga bisa gunakan method lainnya di https://go.dev/src/net/http/method.go
// }
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request) {
        // Read request Method
        method := r.Method

        // Read URL and parameter value:
        url := r.URL
        urlParam := r.URL.Query().Get("param")

        // Read request Header
        headerValues := r.Header.Values("foo") // menampilkan ["bar1", "bar2"]
        headerGet := r.Header.Get("foo")       // menampilkan "bar1" . Hanya nilai pertama yang dikembalikan.

        // Read request Body
        body, err := io.ReadAll(r.Body)
        if err != nil {
            w.Write([]byte("Error read Body" + err.Error()))
            return
        }

        writer := fmt.Sprintf("Method:\t\t %v \nUrl:\t\t %v \nUrl param:\t %v \nHeaderValues:\t %v \nHeaderGet:\t %v \nBody:\t\t %v", method, url, urlParam, headerValues, headerGet, string(body))
        w.Write([]byte(writer))
    })

    fmt.Println("server started at localhost:8080")
    http.ListenAndServe(":8080", nil)
}