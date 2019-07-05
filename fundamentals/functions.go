package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func difference(a, b int) int {
	return a - b
}

func product(a, b int) int {
	return a * b
}

func quotient(a, b int) int {
	return a / b
}

func multipleReturn() (int, int) {
	return 100, 109
}

func variadicFunc(nums ...int) int {
	fmt.Println(nums) // Slice

	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func main() {

	final := variadicFunc(10, 20, 30, 40, 50)
	defer fmt.Println(final)

	fmt.Println("This is the function script!")

	fmt.Println(sum(1, 2))
	fmt.Println(difference(10, 9))
	fmt.Println(product(100, 5))
	fmt.Println(quotient(24, 4))

	x, y := multipleReturn()
	fmt.Println(x, y)

	_, b := multipleReturn()
	fmt.Println(b)

	sum := variadicFunc(1, 2, 3, 4, 5)

	fmt.Println(sum)

	xi := []int{1, 2, 3, 4, 5}
	sumTwo := variadicFunc(xi...)

	fmt.Println(sumTwo)

	// Annonymous Functions
	func(x int) {
		fmt.Println("Testing annoynmous", x)
	}(100)
}
