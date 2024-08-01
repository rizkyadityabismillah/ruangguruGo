package testing2

import (
	"testing"
	"time"
)

func TestGetAge(t *testing.T){
	var person = Person{
		Name: "Afrida Maharani",
		DateBirth: time.Date(2005, 1,1,0,0,0,0, time.UTC),
	}
	var expected int = 19
	
	if person.GetAge() != expected{
		t.Errorf("Expected %d, but got %d", expected, person.GetAge())
	}
}