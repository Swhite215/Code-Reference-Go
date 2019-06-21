package main

import "fmt"

var x = 100   // Use var outside function body, program scope
var a int     // Zero value = 0
var b float32 // Zero value = 0.0
var c bool    // Zero value = false
var d string  // Zero value = ""

func main() {
	y := 43 // Short declaration operator within function body

	fmt.Println(x + y)
}
