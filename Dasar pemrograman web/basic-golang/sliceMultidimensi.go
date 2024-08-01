package main

import "fmt"

func main(){
	// sliceM := [][]int{{1,2,3,4}, {5,6,7}}

	// for _, v := range sliceM{
	// 	fmt.Println(v[1])
	// }
	// fmt.Println(sliceM)
	//MULTIDIMENSI INT
	s1 := [][]int{
		{1,2},
		{3,4},
		{5,6},
	}
	fmt.Println("MULTIDIMENSI SLICE array = ",s1)
	s2 := [][]string{
		[]string{"Rida", "afrida"},
		[]string{"maharani", "triadi"},
		[]string{"Rizky", "Aditya"},
	}
	fmt.Println("MUltimensi slice string = ", s2)

	for i := 0; i < len(s2); i++{
		fmt.Println("MUltiperulangan = ", s2[i])
	}
	n := len(s2)
	m := len(s2[0])
	for i :=0; i < n ; i++ {
		for j :=0; j <m ; j++{
			fmt.Println(s2[i][j], " ")
		}
		fmt.Println("\n")
	} 
}