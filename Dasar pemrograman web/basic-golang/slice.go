package main

import (
	"fmt" 
	
)

// func main(){
// 	arr := [...]string{"Rida","Afrida","Iky","Rizky","Kroos","modric","Klok"}
// 	fmt.Println(arr)
// 	mySlice := arr[:]
// 	fmt.Println("My slice is : ", mySlice)
// 	arrAppend := append(mySlice, "Dwi")
// 	fmt.Println(arrAppend)
// 	fmt.Println("Menampilkan panjang slice \t", len(mySlice))
// 	fmt.Printf("menampilkan Capacity \t", cap(mySlice))
// }

// func main(){
// 	mySlice1 := make([]string, 2,2)
// 	mySlice1[0] = "Senin"
// 	mySlice1[1] = "Selasa"
// 	//CONCAT
// 	mySlice2 :=  []string{"Rabu","Kamis"}
// 	gabungan1 := append(mySlice1,mySlice2...)
// 	fmt.Println(gabungan1)
// 	mySlice3 := []string {"Jumat","Sabtu"}
// 	gabungan2 := append(gabungan1,mySlice3...)
// 	fmt.Println(gabungan2)
// 	copydays := make([]string,len(mySlice1), (cap(mySlice1)+1)*2) 
// 	copy(copydays, mySlice1)
// 	fmt.Println("days \t\t :", mySlice1)
// 	fmt.Println("cap \t\t :", cap(mySlice1))
// 	fmt.Println("copied days:\t :", copydays)
// 	fmt.Println("Cap of copied days:", cap(copydays))

// 	copydays[1] = "Sabtu"
// 	mySlice1[1] = "Minggu"
// 	fmt.Println("_________")
// 	fmt.Println("days update \t :",mySlice1)
// 	fmt.Println("days update \t :",copydays)

// }

func main() {
	//SLICE MULTIDIMENSI
	sliceMultidimensi := [][]string{
		[]string{"Afrida", "Maharani"},
		[]string{"Rizky", "Aditya"},
	}
	fmt.Println("Slice string :",sliceMultidimensi)
	sliceInt := [][]int{
		{1,2},
		{3,4},
		{6,7},
	}
	fmt.Println("Slice int :", sliceInt)

	for i := 0; i < len(sliceInt); i++{
		fmt.Println(sliceInt[i])
	} 

	n := len(sliceInt)
	m := len(sliceInt[0])
	fmt.Println("rows :",n ,"\ncolumns : " , m)

	for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            fmt.Println(sliceInt[i][j])
        }
        fmt.Print("\n")
    }
}