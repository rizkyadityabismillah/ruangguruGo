package main

import "testing"

func TestArea(t *testing.T){ // fungsi ini digunakan untuk mengetes method Area di struct 'Cube'
	var cube = Cube{side: 2} // menyiapkan data untuk diuji
	var expected float64 = 24.0 //NIlai yang diharapkan
	
	if cube.Area() != expected{ // jika hasil tidak sesuai dengan yang diharapkan, maka test dianggap gagal
		t.Errorf("Expected %f ,but got %f",expected, cube.Area())
	}
}

func TestCircumference(t *testing.T){
	var cube = Cube {side:2}
	var expected float64 = 24.0

	if cube.Circumference() != expected{
		t.Errorf("Expected %f ,but got %f",expected, cube.Circumference())
	}
}

func TestVolume(t *testing.T){
	var cube = Cube{side :2}
	var expected float64 = 8.0

	if cube.Volume() != expected{
		t.Errorf("Expected %f, but got %f", expected, cube.Volume())
	}
}