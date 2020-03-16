package assignmentthree

import "fmt"

type person struct {
	name   string
	gender string
	age    int
}

func NewPerson() *person {
	var p person

	return &p
}

func (p *person) Name(name string) *person {
	p.name = name

	return p
}

func (p *person) Gender(gender string) *person {
	p.gender = gender

	return p
}

func (p *person) Age(age int) *person {
	p.age = age

	return p
}

func main() {
	jon := NewPerson().Name("Jon Snow").Gender("Male").Age(24)
	fmt.Println(jon)
}
