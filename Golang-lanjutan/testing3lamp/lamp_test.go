package testing3lamp

import "testing"

func TestTurnOn(t *testing.T){
	//Arrange bagian di mana kita membuat data yang akan digunakan untuk testing.
	lamp := Lamp {State: false}
	//Act bagian di mana kita menjalankan fungsi yang akan kita test.
	lamp.TurnOn()
	//Assert  bagian di mana kita membandingkan hasil dari fungsi yang kita test dengan hasil yang diharapkan.
	if lamp.State != true{
		t.Errorf("lamp state should be true, got %v", lamp.State)
	}
}
func TestTurnOff(t *testing.T){
	lamp := Lamp {State:  true}
	lamp.TurnOff()
	if lamp.State != false{
		t.Errorf("lamp state should be true, got %v", lamp.State)
	}
}