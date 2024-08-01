package server

import (
    "encoding/json"
    "net/http"

    "github.com/rest/model"
)

func Students(w http.ResponseWriter, r *http.Request) {
	data := []model.Student{
		{Name: "Aditira", Age: 22},
        {Name: "Dito", Age: 24},
        {Name: "Ojan", Age: 30},
		{name: "Afrida", Age: 19},
	}
	jsonInByte, err := json.Marshal(data)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonInByte)

}