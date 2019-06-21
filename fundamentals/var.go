package main

import "fmt"

var x = 100 // Use var outside function body, program scope

func main() {
	y := 43 // Short declaration operator within function body

	fmt.Println(x + y)
}
