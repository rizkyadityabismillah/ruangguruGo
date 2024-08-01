package main

import (
	"fmt"
	
)

// func main(){
// 	students := [3]string{"Afrida", "maharani","sayang"}
// 	fmt.Println(students[0])
// }

//Default nilai tiap elemen array pada awalnya tergantung dari tipe datanya. Jika int maka tiap element zero value-nya adalah 0, jika bool maka false, dan seterusnya
// func main(){
// 	arrNum := [4] int{3,4}
// 	fmt.Println(arrNum)

// 	arrBool := [4] bool{true, false}
// 	fmt.Println(arrBool)
// }

// func main(){
// 	var names [4] string
// 	names[0]= "afrida"
// 	names[1]= "maharani"
// 	names[2]= "pacar"
// 	names[3]= "sayang"
// 	fmt.Println(names)
// }

//CARRA INISIALISASI ARRAY
//INI HORIZONTAL
// func main(){
// 	var fruits = [4]string{"Tea jus", "JasJus","Sprit","Apel"}
// 	fmt.Println("jumllah minuman \t:", len(fruits))
// 	fmt.Println("ini semua minuman\t:", fruits)
// }

// INI VERTIKAL
// func main() {
// 	var fruits = [4]string{
// 		"Tea Jus",
// 		"JasJus",
// 		"Sprit",
// 		"Apel",
// 	}
// 	fmt.Println(len(fruits))
// 	fmt.Println(fruits)
// }

//Tanpa jumlah elemen
// func main(){
// 	var arrFruit = [...] string{"Es Buah", "Es Campur", "Es Jeruk"}
// 	fmt.Println("Jumlah Es :",len(arrFruit))
// 	fmt.Println("Jenis Es :",arrFruit)
// }

//MULTIDIMENSI ARRAY
// YANG ELEMEN ARRAYNYA DIDALAM ARRAY
// func main() {
// 	var students = [2][3] string{{"Rizky","Aditya","Triadi"},[3]string{"Afrida","Maharani","Sayang"}}
// 	var number = [2][3] int{{1,2,3},{4,5,6}}
// 	fmt.Println(students[0][0])
// 	fmt.Println(students[1][0])

// 	fmt.Println(number[1])
// }

//ARRAY LOOPINGNG
// func main(){
// 	fruit := [...]string{"es buah", "es anggur", "es jambu", "es teajus"}

// 	for i := 0 ; i < len(fruit);i++{
// 		fmt.Println(i,fruit[i])
// 	}
// }

//ARRAY FOR - RANGE
// fruits := [...]string{"es buah", "es anggur", "es jambu", "es teajus"}

// for _,fruit := range fruits {
// 	fmt.Println( fruit)
// }
// var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
// var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

// func main() {
// 	idNameStudents :=  strings.Split(Students,",")
// 	for _, value := range idNameStudents{
// 		item := strings.Split(value, "_")
// 		if len(item) == 3  {
// 			id := item[0]
// 			name := item[1]
// 			major := item[2]
// 			fmt.Println(id)
// 			fmt.Println(name)
// 			fmt.Println(major)
// 		}
// 		fmt.Println(item)
// 	}

// }
func main() (data string){
	data = "hello"
	return

}