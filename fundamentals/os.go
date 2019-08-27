package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("top.txt")

	if err != nil {
		fmt.Println("Error: ", err)
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	bs := make([]byte, stat.Size())
	_ , err = file.Read(bs)

	if err != nil {
		fmt.Println("Error: ", err)
	}

	str := string(bs)
	fmt.Println(str)


}
