package main

import "fmt"

func main() {

	// Array
	var x [5]int
	x[3] = 1000

	fmt.Println(x[3])
	fmt.Println(len(x))

	b := [3]int{0, 1, 2}

	fmt.Println(b)
	
	
	for index, value := range b {
		fmt.Println(index, value)
	}
	
	// Example where change to slice affects underlying array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
