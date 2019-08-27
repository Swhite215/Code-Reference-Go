package main

import (
	"os"
)

func main() {
	file, err := os.Create("topTwo.txt")
	if err != nil {
		return
	}

	defer file.Close()

	file.WriteString("Test")
}
