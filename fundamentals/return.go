package main

import (
	"fmt"
	"strings"
)

func returnUpper(firstName, lastName string) (upperFirst, upperLast string) {
	upperFirst = strings.ToUpper(firstName)
	upperLast = strings.ToUpper(lastName)

	return
} 

func main() {
	fmt.Println(returnUpper("Spencer", "White"))
}
