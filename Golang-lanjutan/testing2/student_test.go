package testing2

import (
    "testing"
    "time"
)

func TestGetPredicate(t *testing.T) {
    var student = Student{
        Person: Person{
            Name: "John Doe",
            DateBirth: time.Date(1980, 1, 1, 0, 0, 0, 0, time.UTC),
        },
        Batch: "Batch 1",
        Grade: 90,
    }

    var expected string = "John Doe got Predicate: Excellent"
    if student.GetPredicate() != expected {
        t.Errorf("Expected %s, but got %s", expected, student.GetPredicate())
    }
}