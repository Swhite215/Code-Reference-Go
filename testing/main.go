package main

import (
	"fmt"
	"testing/utilities" //Works if utilities was in a separate folder and declared a package
)

func main() {
	fmt.Println("Hello!")

	fmt.Println(utilities.Sum(10,10))
	fmt.Println(utilities.Difference(10,10))
	fmt.Println(utilities.Product(10,10))
	fmt.Println(utilities.Quotient(10,10))
}
