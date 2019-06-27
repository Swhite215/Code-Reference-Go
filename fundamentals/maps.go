package main

import "fmt"

func main() {

	m := map[string]int{
		"Person 1": 23,
		"Person 2": 45,
	}

	fmt.Println(m)

	tools := make(map[int]string)

	tools[1] = "Saw"
	tools[2] = "Drill"

	fmt.Println(tools)

	v, ok := tools[3]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := tools[3]; ok {
		fmt.Println("Value does exists and is okay: " + v)
	} else {
		fmt.Println("Value does not exist and is not okay:")
	}

	delete(tools, 1)
	fmt.Println(tools)

	tools[1] = "Screwdriver"

	for k, v := range tools {
		fmt.Println(k, v)
	}
}
