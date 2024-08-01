package multipledata4
import "testing"
// func SumTwoNumber(a int, b int) int{
// 	return a + b
// }
var person *Person = NewPerson("John Doe", 20, 100_000)
func SetupTest(t *testing.T) {
    t.Log("Setup test case")

    person.SetBonus(12)
}

func TestAddIncome(t *testing.T) {
    // menjalankan setup test
    SetupTest(t)

    person.AddIncome(100000)
    if person.GetWealth() != 150000.0 {
        t.Errorf("Expected %f, but got %f", 150000.0, person.GetWealth())
    }
}

func TestAddCost(t *testing.T) {
    // menjalankan setup test
    SetupTest(t)

    person.AddCost(10000)

    if person.GetWealth() != 190000.0 {
        t.Errorf("Expected %f, but got %f", 190000.0, person.GetWealth())
    }
}
// setup test untuk mempersiapkan data struct 'Person'

