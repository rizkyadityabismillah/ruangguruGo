package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	// Jika id kosong atau name kosong
	if id == "" || name == "" {
		// Return string
		return "ID or Name is undefined!"
	}

	// Jika id tidak sama dengan 5
	if len(id) != 5 {
		// Return string
		return "ID must be 5 characters long!"
	}

	// Output: A1234_Aditira_TI B2131_Dito_TK A3455_Afis_MI
	studentData := strings.Split(Students, ", ")

	// Output:
	// [0] = A1234_Aditira_TI
	// [1] = B2131_Dito_TK
	// [2] = A3455_Afis_MI
	for _, student := range studentData {
		// Jika [0] = A1234_Aditira_TI
		// [0] = A1234
		// [1] = Aditira
		// [2] = TI
		info := strings.Split(student, "_")

		// Jika A123 sama dengan id dan Aditra sama dengan name
		if info[0] == id && info[1] == name {
			// Return string dan nilai dari name, major
			return fmt.Sprintf("Login berhasil: %s (%s)", name, info[2])
		}
	}

	// Jika data tidak tersedia
	return "Login gagal: data mahasiswa tidak ditemukan"

}

func Register(id string, name string, major string) string {
	// Jika id kosong atau name kosong atau major kosong
	if id == "" || name == "" || major == "" {
		// Return string
		return "ID, Name or Major is undefined!"
	}

	// Jika panjang id tidak sama dengan 5
	if len(id) != 5 {
		// return string
		return "ID must be 5 characters long!"
	}

	// Output: A1234_Aditira_TI B2131_Dito_TK A3455_Afis_MI
	studentData := strings.Split(Students, ", ")

	// Output:
	// [0] = A1234_Aditira_TI
	// [1] = B2131_Dito_TK
	// [2] = A3455_Afis_MI
	for _, student := range studentData {
		// [0] = A1234
		// [1] = Aditira
		// [2] = TI
		info := strings.Split(student, "_")

		// Jika A1234 == id yang diinput
		if info[0] == id {
			// return string
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	// Menambahkan data baru kedalam variable students
	// ID_NAME_MAJOR yang diinput
	Students += fmt.Sprintf(", %s_%s_%s", id, name, major)

	// Mengembalikan nilai dari NAME MAJOR
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {
	// Jika kode kosong
	if code == "" {
		// return string
		return "Code is undefined!"
	}

	// Output: TI_Teknik Informatika TK_Teknik Komputer SI_Sistem Informasi MI_Manajemen Informasi
	programData := strings.Split(StudentStudyPrograms, ", ")

	// Output:
	// [0]TI_Teknik Informatika
	// [1]TK_Teknik Komputer
	// [2]SI_Sistem Informasi
	// [3]MI_Manajemen Informasi
	for _, program := range programData {
		// Jika info[0] == [0]TI_Teknik Informatika
		// Output:
		// [0] TI
		// [1] Teknik Informatika
		info := strings.Split(program, "_")

		// Jika TI == code jurusan yang diinput (TI)
		if info[0] == code {
			// return Teknik Informatika
			return info[1]
		}
	}
	// Return String
	return "Program studi tidak ditemukan"
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
