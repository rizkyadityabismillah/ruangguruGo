package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Study struct{
	StudyName string   `json:"study_name"`
	StudyCredit int     `json:"study_credit"`
	Grade string   `json:"grade"`
}
type Report struct {
	Id string  `json:"id"`
	Name string    `json:"name"`
	Date string `json:"date"`
	Semester  int `json:"semester"`
	Studies    []Study `json:"studies"`
}


// gunakan fungsi ini untuk mengambil data dari file json
// kembalian berupa struct 'Report' dan error
func ReadJSON(filename string) (Report, error) {
	f,err := os.Open(filename) // buka file json
	if err != nil { 
		return Report{},err
}
	defer f.Close()           
	byteData, err := io.ReadAll(f)
	if err != nil{
		return Report{},err
	}

	var report Report
	err =json.Unmarshal(byteData,&report)

	if err != nil{
		return Report{},err
	}
	return report,nil
}

func GradePoint(report Report) float64 {
	if len(report.Studies) == 0{
		return 0.0
	}

	var accumIp  float64
	var accumCredit  int

	var gradeMap = map[string]float64{
			"A": 4.0,
			"AB":3.5,
			"B":3.0,
			"BC":2.5,
			"C":2.0,
			"CD":1.5,
			"D":1.0,
			"DE":0.5,
			"E":0,
		}
		for _ ,study := range report.Studies{
		accumIp += float64(study.StudyCredit) * gradeMap[study.Grade]
		accumCredit += study.StudyCredit
		}
	return accumIp / float64(accumCredit) 
}

func main() {
	// bisa digunakan untuk menguji test case
	report, err := ReadJSON("report.json")
	if err != nil {
		panic(err)
	}

	gradePoint := GradePoint(report)
	fmt.Println(gradePoint)
}
