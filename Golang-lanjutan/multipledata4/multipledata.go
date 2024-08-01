package multipledata4


// type TestCase struct{
// 	a int
// 	b int
// 	expected int
// }
// func TestNewNumber(t *testing.T){
// var listcase = []TestCase{ // menyiapkan semua data test case
// 	{a: 2, b:3 , expected:5},
// 	{a: -10, b:15 , expected:-5},
// 	{a: 5, b: -7, expected: -2 },
// 	{a: -5, b:-5, expected: -10},
// }
// for _, test := range listcase { // melakukan perulangan untuk menjalankan semua test case
// 	result := SumTwoNumber(test.a, test.b)
// 	if result != test.expected {
// 		t.Errorf("SumTwoNumber%d,Expected %d" , result , test.expected )
// }
// }

type Person struct{
	Name string
	Age int
	BaseSalary float64
	Wealth float64
}

func NewPerson(name string, age int, baseSalary float64) *Person{
	return &Person{
		Name:  name,
		Age:   age,
		BaseSalary : baseSalary,
	}
}

func(p *Person) GetWealth()float64{
	return p.Wealth
}

func(p *Person) SetBonus(MonthOfWork int){
	p.Wealth += (float64(MonthOfWork)/12) * p.BaseSalary * 0.5
}

func (p *Person) AddIncome(Income float32){
	p.Wealth += float64(Income)
}
func (p *Person) AddCost(Cose float64){
	p.Wealth -= Cose
}

func (p *Person) ClearWealth(){
	p.Wealth = 0
}