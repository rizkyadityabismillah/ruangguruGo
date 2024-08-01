package main

import (
	"a21hc3NpZ25tZW50/helper"
	"fmt"
	"strings"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	
	if id == "" ||  name == ""{
		 return "ID or Name is undefined!"
	}
	if len(id) != 5 {
		return "ID must be 5 characters long!"
	} 
	
	for _, student := range strings.Split(Students, ", ") {
			// Output: A1234_Aditira_TI B2131_Dito_TK A3455_Afis_MI
		item := strings.Split(student, "_")
		// [0] = A1234
		// [1] = Aditira
		// [2] = TI
		if item[0] == id && item[1] == name {
			return fmt.Sprintf("Login berhasil: %s (%s)", name, item[2])
		}
	}
		return "Login gagal: data mahasiswa tidak ditemukan"
}
	
func Register(id string, name string, major string) string {
	// Jika id kosong atau name kosong atau major kosong
	if id == "" || name == "" || major == ""{
		return "ID, Name or Major is undefined!"
	}
	//JIKA ID TIDAK SAMA DENGAN 5
	if len(id) != 5{
		return "ID must be 5 characters long!"
	}
	for _, student := range strings.Split(Students, ", "){
		item := strings.Split(student, "_")
		if item[0] == id{
			return "Registrasi gagal: id sudah digunakan"
			} 
		} 
			Students += fmt.Sprintf(", %s_%s_%s", id, name, major)
			return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
			
}

func GetStudyProgram(code string) string {
	var messageCodeMK string
	if code == ""{
		messageCodeMK = "Code is undefined!"
	}
	for _, studyCodemk := range strings.Split(StudentStudyPrograms, ", "){
		studyProgram := strings.Split(studyCodemk,"_")
	// Output:
	// [0]TI_Teknik Informatika
	// [1]TK_Teknik Komputer
	// [2]SI_Sistem Informasi
	// [3]MI_Manajemen Informasi
		if  studyProgram[0] == code{
			messageCodeMK = studyProgram[1]
			break
	}
}
return messageCodeMK // TODO: replace this
}
func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
