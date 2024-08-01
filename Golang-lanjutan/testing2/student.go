package testing2

type Student struct{
	Person
	Batch string
	Grade int
}

func (s Student) GetPredicate() string {
    switch {
    case s.Grade == 100:
        return s.Name + " got Predicate: Perfect"
    case s.Grade < 100 && s.Grade >= 90:
        return s.Name + " got Predicate: Excellent"
    case s.Grade < 90 && s.Grade >= 80:
        return s.Name + " got Predicate: Good"
    case s.Grade < 80 && s.Grade >= 70:
        return s.Name + " got Predicate: Acceptable"
    default:
        return s.Name + " got Predicate: Bad"
    }
}