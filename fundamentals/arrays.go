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
}
