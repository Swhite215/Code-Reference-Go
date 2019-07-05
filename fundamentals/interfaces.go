package main

import "fmt"

type secretAgent struct {
	firstName string
	lastName  string
	age       int
}

type researcher struct {
	firstName string
	lastName  string
	subject   string
}

type human interface {
	speak()
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.firstName, s.lastName)
}

func (r researcher) speak() {
	fmt.Println("I am", r.firstName, r.lastName, " and I research", r.subject)
}

func identity(h human) {
	fmt.Println("I am a human", h)
}

func main() {

	sa1 := secretAgent{
		firstName: "James",
		lastName:  "Bond",
		age:       32,
	}

	r1 := researcher{
		firstName: "Albert",
		lastName:  "Einstein",
		subject:   "Physics",
	}

	sa1.speak()
	r1.speak()

	identity(sa1)
	identity(r1)

}
