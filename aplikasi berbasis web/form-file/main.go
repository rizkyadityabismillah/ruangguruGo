package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Handler route / isinya proses untuk menampilkan landing page (file view.html). Method yang diperbolehkan mengakses rute ini hanya GET.
func routeIndexGet(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET" {
	http.Error(w, "", http.StatusBadRequest)
	return
}
	var tmpl = template.Must(template.ParseFiles("view.html"))
	var err = tmpl.Execute(w, nil)

	if err != nil{
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
}
// Selanjutnya siapkan handler untuk rute /process, yaitu fungsi routeSubmitPost. Gunakan statement r.ParseMultipartForm(1024) untuk parsing form data yang di kirim.
func routeSubmitPost(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "", http.StatusBadRequest)
		return
	}
	// Method ParseMultipartForm() digunakan untuk mem-parsing form data. Argumen 1024 pada method tersebut adalah maxMemory
	if err := r.ParseMultipartForm(1024); err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	alias := r.FormValue("alias")

uploadedFile, handler, err := r.FormFile("file")
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
defer uploadedFile.Close()
//mengambil relative path dari proyek
dir, err := os.Getwd()
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

//membuat nama file baru yang akan disimpan
filename := handler.Filename
if alias != "" { //kalau alias disi maka nama file = alias
    filename = fmt.Sprintf("%s%s", alias, filepath.Ext(handler.Filename))
}

//membentuk lokasi tempat menyimpan file
fileLocation := filepath.Join(dir, "images", filename)
targetFile, err := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}
defer targetFile.Close()

//mengisi file baru dengan data dari file yang ter-upload
if _, err := io.Copy(targetFile, uploadedFile); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
}

w.Write([]byte("done"))
}

func HandlerImage(w http.ResponseWriter, r *http.Request){
	imageName := r.URL.Query().Get("image-name") // mengambil nama image dari query url
	fileByte, err:= ioutil.ReadFile("./images/" + imageName) // membaca file image menjadi bytes
	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("filenotfound"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileByte)
}
func main(){
	http.HandleFunc("/", routeIndexGet) /// digunakan untuk landing page.
	http.HandleFunc("/process", routeSubmitPost) ///process digunakan ketika proses upload selesai.
	http.HandleFunc("/view", HandlerImage) ///view?image-name=<nama image beserta extension nya> digunakan untuk menampilkan gambar yang ada di server.

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe( ":9000", nil )
}