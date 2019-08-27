package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("top.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}

	str := string(bs)
	fmt.Println(bs)
	fmt.Println(str)
}
