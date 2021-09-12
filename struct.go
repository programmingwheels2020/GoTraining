package main

import "fmt"

type Person struct {
	name     string
	age      int
	location string
}

type Employee struct {
	person      Person
	designation string
}

func main() {
	/*p := Person{
		name:     "Lenin Lawrence",
		age:      70,
		location: "Kochin",
	}

	/*fmt.Println(p.age)*/
	emp := Employee{
		person: Person{
			"Lenin", 70, "Kochin",
		},
		designation: "SE",
	}

	fmt.Println(emp.person.name)
}
