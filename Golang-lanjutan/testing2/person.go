package testing2

import "time"

type Person struct {
	Name      string
	DateBirth time.Time
}
func (p Person) GetAge() int{
	return time.Now().Year() - p.DateBirth.Year()
}
func (p Person) GetName() string{
	return p.Name
}