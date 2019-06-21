package main

import "fmt"

// Entry point to the program, when code exits main, program is over
func main() {
	fmt.Println("Hello World!")

	n, err := fmt.Println("Hello", 100, true)

	fmt.Println(n)   // Returns length of arguments added together
	fmt.Println(err) // Returns error
}
