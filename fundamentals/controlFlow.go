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

	//Conditional - #1
	a := 2
	if a == 2 {
		fmt.Println("A is 2!")
	}

	//Conditional - #2
	b := true
	if b == false {
		fmt.Println("B is false!");
	} else {
		fmt.Println("B is true!")
	}

	//Conditional - #3
	c := "Ten"
	if c == "One" {
		fmt.Println("B is One!");
	} else if c == "Two" {
		fmt.Println("B is Two!")
	} else {
		fmt.Println("B is not One or Two!")
	}

	///Conditional - #4
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
	
}

// Sequential, Iterative, Conditional
