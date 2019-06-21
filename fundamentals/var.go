package main

import "fmt"

var x = 100   // Use var outside function body, program scope
var a int     // Zero value = 0
var b float32 // Zero value = 0.0
var c bool    // Zero value = false
var d string  // Zero value = ""

type newType int

var e newType = 100

func main() {
	y := 43 // Short declaration operator within function body

	fmt.Println(x + y)

	// Printing types
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)

	// Conversion
	a = int(e)
	fmt.Printf("%T\n", a)

}
