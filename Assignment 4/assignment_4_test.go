package assignmentfour

import (
	"testing"
)

func FirstCase(t *testing.T) {
	person := Person{Name: "Agus", Gender: "Male", Age: 20}

	err := person.Validate()
	if err != nil {
		t.Errorf("It should not return error")
	}
}

func SecondCase(t *testing.T) {
	person := Person{Name: "", Gender: "Male", Age: 20}

	err := person.Validate()
	if err == nil {
		t.Errorf("It should return error")
	}

	if err.Error() != "Name cannot be empty" {
		t.Errorf("It should return error with message " + "\"Name cannot be empty\"")
	}
}

func ThirdCase(t *testing.T) {
	person := Person{Name: "Agus", Gender: "Shemale", Age: 20}

	err := person.Validate()
	if err == nil {
		t.Errorf("It should return error")
	}

	if err.Error() != "Gender is either Male or Female" {
		t.Errorf("It should return error with message " + "\"Gender is either Male or Female\"")
	}
}

func FourthCase(t *testing.T) {
	person := Person{Name: "Agus", Gender: "Male", Age: -20}

	err := person.Validate()
	if err == nil {
		t.Errorf("It should return error")
	}

	if err.Error() != "There is no such thing as negative age" {
		t.Errorf("It should return error with message " + "\"There is no such thing as negative age\"")
	}
}

func TestValidate(t *testing.T) {
	t.Run("Test data complete", FirstCase)
	t.Run("Test data with 'Name' empty", SecondCase)
	t.Run("Test data with 'Gender' not Male or Female", ThirdCase)
	t.Run("Test data with 'Age' negative value", FourthCase)
}
