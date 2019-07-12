package main

import "fmt"

// Argument passed by value, will not change the passed value in main
func zeroval(ival int) {
	ival = 0
}

// Assigns a value to a dereferenced pointer, changing the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)        // Address in Memory
	fmt.Printf("%T\n", a)  // Int
	fmt.Printf("%T\n", &a) // Pointer to Int
	fmt.Println(*&a)       // Value at Address in Memory

	// When to use pointers, passing large chunk of data, change a value at a certain address

	i := 1
	fmt.Println("initial:", i)
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
