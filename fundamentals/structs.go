package main

import "fmt"

type person struct{
	firstName string
	lastName string
	age int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {

	personOne := person{
		firstName: "Person",
		lastName: "One",
		age: 54,
	}

	fmt.Println(personOne)

	personTwo := person{
		firstName: "Person",
		lastName: "Two",
		age: 12,
	}

	fmt.Println(personTwo)

	bond := secretAgent{
		person: person{
			firstName: "James",
			lastName: "Bond",
			age: 32,
		},
		licenseToKill: true,
	}

	fmt.Println(bond)
	fmt.Println(bond.firstName)

	people := []person{personOne, personTwo}

	fmt.Println(people)

	// Annonymous Struct
	annonymous := struct {
		firstName string
		lastName string
		age int
	}{
		firstName: "Annonymouse",
		lastName: "Declaration",
		age: 100000,
	}

	fmt.Println(annonymous)
}
