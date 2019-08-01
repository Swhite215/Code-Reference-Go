package main

import (
	"fmt"
	"encoding/json"
)

type ColorGroup struct {
	ID int
	Name string
	Colors []string
}

type Animal struct {
	Name string
	Order string
}

type person struct {
	First string
	Last string
	Age int
}

func main() {

	group := ColorGroup{
		ID: 1,
		Name: "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}

	b, err := json.Marshal(group) // Takes a slice of structs and returns a byte string
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(b)

	var jsonBlob = []byte(`[{"Name": "Platypus", "Order": "Monotremata"}]`)

	var animals []Animal
	err2 := json.Unmarshal(jsonBlob, &animals) // Takes a byte string and returns a slice of structs

	if err != nil {
		fmt.Println("error", err2)
	}

	fmt.Printf("%+v", animals)

	p1 := person{
		First: "James",
		Last: "Bond",
		Age: 32,
	}

	p2 := person{
		First: "Idris",
		Last: "Elba",
		Age: 46,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people) // Takes a slice of structs and returns a byte string

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(string(bs)) // Returns empty objects if struct fields are lowercase

	var peopleTwo []person
	err3 := json.Unmarshal(bs, &peopleTwo)

	if err3 != nil {
		fmt.Println("error:", err3)
	}

	fmt.Println(peopleTwo)
}
