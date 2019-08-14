package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 4 << uint(i)
		if pow[i] >= 16 {
			break
		}
	}

	fmt.Println(pow)

	pow = make([]int, 10)
	for i := range pow {
		if i % 2 == 0 {
			continue
		}

		fmt.Println(i)
		pow[i] = 1 << uint(i)
	}

	fmt.Println(pow)
}
