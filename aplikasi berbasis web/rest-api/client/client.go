package client

import (
    "encoding/json"
    "html/template"
    "io/ioutil"
    "net/http"
    "path"

    "github.com/rest/model"
)

func Student(w http.ResponseWriter, r *http.Request) {
    var filepath = path.Join("template", "student.html")

    var tmpl, err = template.ParseFiles(filepath)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Request data student dari client ke server
    var client = &http.Client{}
    req, err := http.NewRequest("GET", "http://localhost:8080/api/student", nil)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
	resp, err := client.Do(req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    // Membaca response body dari server
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    var student []model.Student
    err = json.Unmarshal(body, &student)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }

    var data = map[string]interface{}{
        "title": "Student Bootcamp",
        "data":  student,
    }

    // Render data ke template
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}