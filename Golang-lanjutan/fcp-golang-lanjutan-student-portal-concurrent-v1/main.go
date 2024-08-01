package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"a21hc3NpZ25tZW50/helper"
	"a21hc3NpZ25tZW50/model"
)

type StudentManager interface {
	Login(id string, name string) error
	Register(id string, name string, studyProgram string) error
	GetStudyProgram(code string) (string, error)
	ModifyStudent(name string, fn model.StudentModifier) error
}

type InMemoryStudentManager struct {
	sync.Mutex
	students             []model.Student
	studentStudyPrograms map[string]string
	//add map for tracking login attempts here
	failedLoginAttempts  map[string]int
	// TODO: answer here
}

func NewInMemoryStudentManager() *InMemoryStudentManager {
	return &InMemoryStudentManager{
		students: []model.Student{
			{
				ID:           "A12345",
				Name:         "Aditira",
				StudyProgram: "TI",
			},
			{
				ID:           "B21313",
				Name:         "Dito",
				StudyProgram: "TK",
			},
			{
				ID:           "A34555",
				Name:         "Afis",
				StudyProgram: "MI",
			},
		},
		studentStudyPrograms: map[string]string{
			"TI": "Teknik Informatika",
			"TK": "Teknik Komputer",
			"SI": "Sistem Informasi",
			"MI": "Manajemen Informasi",
		},
		//inisialisasi failedLoginAttempts di sini:
		failedLoginAttempts: make(map[string]int),
		// TODO: answer here
	}
}

func ReadStudentsFromCSV(filename string) ([]model.Student, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3 // ID, Name and StudyProgram

	var students []model.Student
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		student := model.Student{
			ID:           record[0],
			Name:         record[1],
			StudyProgram: record[2],
		}
		students = append(students, student)
	}
	return students, nil  
}

func (sm *InMemoryStudentManager) GetStudents() []model.Student {
	return sm.students
}

func (sm *InMemoryStudentManager) Login(id string, name string) (string, error) {
	if id == "" ||  name == ""{
		return"", fmt.Errorf("%s","ID or Name is undefined!")
   }
   if attempts, ok := sm.failedLoginAttempts[id]; ok && attempts >= 3 {
	return "", fmt.Errorf("Login gagal: Batas maksimum login terlampaui")
}
   for _, student := range sm.students {
	if student.ID == id && student.Name == name {
		// Reset failed login attempts if any
		delete(sm.failedLoginAttempts, id) // 5. Reset failed login attempts
		program, ok := sm.studentStudyPrograms[student.StudyProgram]
		if !ok {
			return "", fmt.Errorf("%s","Login gagal: Kode program studi tidak valid")
		}
		return fmt.Sprintf("Login berhasil: Selamat datang %s! Kamu terdaftar di program studi: %s", student.Name, program), nil
	}
}

// Increment failed login attempts
sm.failedLoginAttempts[id]++ // 3. Gunakan failedLoginAttempts untuk melacak percobaan login yang gagal
return "", fmt.Errorf("%s","Login gagal: data mahasiswa tidak ditemukan")
}


func (sm *InMemoryStudentManager) RegisterLongProcess() {
	// 30ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) Register(id string, name string, studyProgram string) (string, error) {
	// 30ms delay to simulate slow processing. DO NOT REMOVE THIS LINE
	sm.RegisterLongProcess()
	if id == "" || name == "" || studyProgram == ""{
		return "", fmt.Errorf("%s", "ID, Name or StudyProgram is undefined!")
	}
	if _, ok := sm.studentStudyPrograms[studyProgram]; !ok {
		return "", fmt.Errorf("%s","Study program " + studyProgram + " is not found")
	}
	for _, student := range sm.students {
		if student.ID == id {
			return "", fmt.Errorf("%s","Registrasi gagal: id sudah digunakan")
		}
	}
	
	newStudent := model.Student{
			ID:          id,
			Name:        name,
			StudyProgram: studyProgram,
		}
		sm.students = append(sm.students, newStudent)
		return fmt.Sprintf("Registrasi berhasil: %s %s", name, "("+studyProgram+")"),nil
}

func (sm *InMemoryStudentManager) GetStudyProgram(code string) (string, error) {
	//Jika kode kosong
	if code == "" {
		return "", fmt.Errorf("%s","Code is undefined!")
	}

	program, available := sm.studentStudyPrograms[code]
	// Dilakukan pengecekan apakah code terdapat dalam map studentStudyPrograms. jika tidak ditemukan, fungsi mengembalikan pesan error
	if !available {
		return "", fmt.Errorf("%s","Kode program studi tidak ditemukan")
	}
	sm.Lock()
	defer sm.Unlock()
	return program, nil
}

func (sm *InMemoryStudentManager) ModifyStudent(name string, fn model.StudentModifier) (string, error) {
	for _, student := range sm.students {
		// [0] = A1234
	    // [1] = Aditira
	    // [2] = TI
		if student.Name == name {
            err := fn(&student)
            if err != nil {
                return "", err // Jika terjadi error saat modifikasi, kembalikan error tersebut
            }
            return "Program studi mahasiswa berhasil diubah.", nil // Jika modifikasi berhasil, kembalikan pesan sukses
        }
    }
	return "", fmt.Errorf("%s","Mahasiswa tidak ditemukan.")
}


func (sm *InMemoryStudentManager) ChangeStudyProgram(programStudi string) model.StudentModifier {
	return func(s *model.Student) error {
		
		// Validasi apakah kode program studi ada di dalam map studentStudyPrograms
		if _, ok := sm.studentStudyPrograms[programStudi]; !ok {
			return fmt.Errorf("%s","Kode program studi tidak ditemukan")
		}
		s.StudyProgram = sm.studentStudyPrograms[programStudi]
		return nil
		} // TODO: replace this
	}

	func (sm *InMemoryStudentManager) ImportStudents(filenames []string) error {
		var wg sync.WaitGroup
		var mu sync.Mutex // Mutex untuk melindungi slice students
		var students []model.Student
	
		// Fungsi untuk impor data mahasiswa dari satu file CSV
		importFromFile := func(filename string) {
			defer wg.Done()
			studentsFromFile, err := ReadStudentsFromCSV(filename)
			if err != nil {
				// Lakukan handling error di sini, misalnya mencetak pesan atau mengembalikan error
				fmt.Printf("Error importing students from %s: %v\n", filename, err)
				return
			}
			// Mengunci slice students sebelum penambahan data
			mu.Lock()
			students = append(students, studentsFromFile...)
			mu.Unlock()
		}
	
		// Membuat goroutine untuk setiap file CSV
		for _, filename := range filenames {
			wg.Add(1)
			go importFromFile(filename)
		}
	
		wg.Wait() // Menunggu semua goroutine selesai
	
		// Mengimpor data mahasiswa ke dalam sistem
		for _, student := range students {
			if msg, err := sm.Register(student.ID, student.Name, student.StudyProgram); err != nil {
				// Lakukan handling error di sini, misalnya mencetak pesan atau mengembalikan error
				fmt.Printf("Error registering student %s: %v\n", student.Name, err)
				// Tetapi kita tidak menghentikan proses jika ada kesalahan
			} else {
				fmt.Println(msg)
				// Kita tidak menggunakan break di sini agar semua mahasiswa diimpor
				break 
			}
		}
	
		return nil
	}
	

func (sm *InMemoryStudentManager) SubmitAssignmentLongProcess() {
	// 3000ms delay to simulate slow processing
	time.Sleep(30 * time.Millisecond)
}

func (sm *InMemoryStudentManager) SubmitAssignments(numAssignments int) {

	start := time.Now()	
	// TODO: answer here
	var wg sync.WaitGroup
	queue := make(chan int, numAssignments)

	// Worker goroutine
	worker := func(id int) {
		defer wg.Done()
		for assignmentID := range queue {
			fmt.Printf("Worker %d: Processing assignment %d\n", id, assignmentID)
			sm.SubmitAssignmentLongProcess()
		}
	}
	// Create worker goroutines
	for i := 0; i < 3000; i++ {
		wg.Add(1)
		go worker(i + 1)
	}

	// Add assignments to the job queue
	for i := 1; i <= numAssignments; i++ {
		queue <- i
	}
	close(queue)
	// Wait for all workers to finish
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Submitting %d assignments took %s\n", numAssignments, elapsed)
}

func main() {
	manager := NewInMemoryStudentManager()

	for {
		helper.ClearScreen()
		students := manager.GetStudents()
		for _, student := range students {
			fmt.Printf("ID: %s\n", student.ID)
			fmt.Printf("Name: %s\n", student.Name)
			fmt.Printf("Study Program: %s\n", student.StudyProgram)
			fmt.Println()
		}

		fmt.Println("Selamat datang di Student Portal!")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Study Program")
		fmt.Println("4. Modify Student")
		fmt.Println("5. Bulk Import Student")
		fmt.Println("6. Submit assignment")
		fmt.Println("7. Exit")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Pilih menu: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			helper.ClearScreen()
			fmt.Println("=== Login ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			msg, err := manager.Login(id, name)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "2":
			helper.ClearScreen()
			fmt.Println("=== Register ===")
			fmt.Print("ID: ")
			id, _ := reader.ReadString('\n')
			id = strings.TrimSpace(id)

			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Study Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.Register(id, name, code)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "3":
			helper.ClearScreen()
			fmt.Println("=== Get Study Program ===")
			fmt.Print("Program Code (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			if studyProgram, err := manager.GetStudyProgram(code); err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Printf("Program Studi: %s\n", studyProgram)
			}
			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "4":
			helper.ClearScreen()
			fmt.Println("=== Modify Student ===")
			fmt.Print("Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Print("Program Studi Baru (TI/TK/SI/MI): ")
			code, _ := reader.ReadString('\n')
			code = strings.TrimSpace(code)

			msg, err := manager.ModifyStudent(name, manager.ChangeStudyProgram(code))
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			fmt.Println(msg)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "5":
			helper.ClearScreen()
			fmt.Println("=== Bulk Import Student ===")

			// Define the list of CSV file names
			csvFiles := []string{"students1.csv", "students2.csv", "students3.csv"}

			err := manager.ImportStudents(csvFiles)
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			} else {
				fmt.Println("Import successful!")
			}

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')

		case "6":
			helper.ClearScreen()
			fmt.Println("=== Submit Assignment ===")

			// Enter how many assignments you want to submit
			fmt.Print("Enter the number of assignments you want to submit: ")
			numAssignments, _ := reader.ReadString('\n')

			// Convert the input to an integer
			numAssignments = strings.TrimSpace(numAssignments)
			numAssignmentsInt, err := strconv.Atoi(numAssignments)

			if err != nil {
				fmt.Println("Error: Please enter a valid number")
			}

			manager.SubmitAssignments(numAssignmentsInt)

			// Wait until the user presses any key
			fmt.Println("Press any key to continue...")
			reader.ReadString('\n')
		case "7":
			helper.ClearScreen()
			fmt.Println("Goodbye!")
			return
		default:
			helper.ClearScreen()
			fmt.Println("Pilihan tidak valid!")
			helper.Delay(5)
		}

		fmt.Println()
	}
}
