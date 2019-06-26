package main

import "fmt"

func main() {
	//Iterative - #1
	for i := 0; i <= 10; i++ {
		fmt.Println(i);
	}

	for i := 0; i <= 10; i++ {
		if i % 2 == 0 {
			continue;
		}
		fmt.Println(i);
	}

	//Iterative - #2
	i := 0
	for i <= 3 {
		fmt.Println(i);
		i++;
	}

	//Iterative - #3
	for {
        fmt.Println("loop")
        break
	}
	
}

// Sequential, Iterative, Conditional
