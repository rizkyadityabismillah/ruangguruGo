package main

import "fmt"

func main (){
	/*var name1, name2 = "john", "rida"
	var result = name1 == name2
	var result2 = 10 > 4

	fmt.Println(result)
	fmt.Println(result2)*/

	//BOOLEAN LOGIC OPERATOR
	
	/*
	&&
	T T = T
	F F = F
	T F = F
	F F = F
	||
	T T = T
	T F = T
	F T = T
	F F = F
	!
	T = F
	F = T
	NB = T = TRUE F = FALSE
	var score, attedance = 90, 70
	var graduated = score > 80 && attedance > 80
	fmt.Println(graduated)*/
	/*var score, attedance = 90, 70
	var graduated = score > 80 || attedance > 80
	fmt.Println(graduated)*/

	//IF
	/*var pacar = "afrida"
	if pacar == "afrida"{
		fmt.Println(pacar)
	}else{
		fmt.Println("hore")
	}*/

	//IF WITH SHORT STATMENT
	/* if x := 80; x >= 80 {
		fmt.Println("selamat anda lulus")
		} else {
			fmt.Println("maaf anda tidak lulus")
			}  */
	//IF WITH SHORT STATMENT variabel sebagai kondisi

	/*if x,y := 70 , 70; x < 50 && y < 50 {
		fmt.Println("Selamat anda lulus")
		fmt.Println("nilai rata rata anda", (x+y)/2)
	} else{
		fmt.Println("Anda tidak lulus")
	} */

	//IF ELSE IF
	/*var score = 95;

	if score == 100{
		fmt.Println("perfect")
	}else if score >= 90 && score <= 100{
		fmt.Println("great job")
	}else if score >= 80 && score <= 90{
		fmt.Println("good")
	} else if score >= 70 && score <= 80 {
		fmt.Println("oke")
	} else{
		fmt.Println("jelek")
	} */
	//SWITH CASE
	// day := 3
	/*switch day {
	case 1 :
		fmt.Println("Minggu")
	case 2 :
		fmt.Println("Senin")
	case 3 :
		fmt.Println("Selasa")
	case 4 :
		fmt.Println("Rabu")
	default:
		fmt.Println("Hari libur")
	}*/

	//beberapa kondisi di dalam 1 case
	// switch day {
	// case 1,2,3,4,5:
	// 	fmt.Println("Weekdays")
	// case 6,7:
	// 	fmt.Println("weekend")
	// default:
	// 	fmt.Println("Kiamat")
	// }

	//CASE CONDITIONAL

	score :=89

	switch{
	case score > 90:
		fmt.Println("Great Job!")
		fallthrough
	case score > 80:
		fmt.Println( "Good Work!")
		fallthrough
	case score <= 70:
		fmt.Println("Not Bad!")
	default:
		fmt.Println("Keep Going...")
	}
}