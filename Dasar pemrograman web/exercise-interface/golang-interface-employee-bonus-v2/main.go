package main

import "fmt"

type Employee interface {
	GetBonus() float64
}

type Junior struct {
	Name         string
	BaseSalary   int
	WorkingMonth int
}

func (j Junior) GetBonus() float64 {
	return 1 * float64(j.BaseSalary) * float64(j.WorkingMonth) / 12
}

type Senior struct {
	Name            string
	BaseSalary      int
	WorkingMonth    int
	PerformanceRate float64
}

func (s Senior) GetBonus() float64 {
	return 2*float64(s.BaseSalary)*float64(s.WorkingMonth)/12 + s.PerformanceRate*float64(s.BaseSalary)
}

type Manager struct {
	Name             string
	BaseSalary       int
	WorkingMonth     int
	PerformanceRate  float64
	BonusManagerRate float64
}

func (m Manager) GetBonus() float64 {
	return 2*float64(m.BaseSalary)*float64(m.WorkingMonth)/12 + (m.PerformanceRate*float64(m.BaseSalary)) + (m.BonusManagerRate*float64(m.BaseSalary))
}

func EmployeeBonus(employee Employee) float64 {
	return employee.GetBonus()
}

func TotalEmployeeBonus(employees []Employee) float64 {
	sum := 0.0
	for _, employee := range employees {
		sum += employee.GetBonus()
	}
	return sum
}
func main() {
	adi := Junior{
		Name:         "Adi",
		BaseSalary:   100000,
		WorkingMonth: 12,
	}
	charlie := Senior{
		Name:            "Adi",
		BaseSalary:      100000,
		WorkingMonth:    12,
		PerformanceRate: 0.5,
	}
	budi := Manager{
		Name:             "Adi",
		BaseSalary:       100000,
		WorkingMonth:     12,
		PerformanceRate:  0.5,
		BonusManagerRate: 0.3,
	}
	fmt.Println(EmployeeBonus(adi))                 // Output
	fmt.Println(EmployeeBonus(charlie))                 // Output
	fmt.Println(EmployeeBonus(budi))                 // Output
}
