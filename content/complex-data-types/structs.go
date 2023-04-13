package main

import "fmt"

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

func main() {
	john := &Person{"John", 34}
	fmt.Println(john)

	robert := NewPerson("Robert", 67)
	fmt.Println(robert)
	myla := NewPerson("Myla", 8)
	persons := []*Person{robert, myla}
	for _, person := range persons {
		fmt.Println(person)
	}
}
