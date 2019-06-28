package main

import "fmt"

type person struct{
	firstName string
	lastName string
	age int
}

type secretAgent struct{
	person
	licenseToKill bool
}

// func (r receiver) identifier(paramters) (returns(s)) { code }

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
}

func main() {
	personOne := secretAgent{
		person: person{
			firstName: "James",
			lastName: "Bond",
			age: 32,
		},
		licenseToKill: true,
	}

	fmt.Println(personOne)
	personOne.speak()
}
