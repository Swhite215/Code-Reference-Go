package main

import "fmt"

func main() {

	// Array
	var x [5]int
	x[3] = 1000

	fmt.Println(x[3])
	fmt.Println(len(x))

	b := [3]int{0, 1, 2}

	fmt.Println(b)

	// Slices - Composite Literal
	y := []int{1, 2, 3, 4, 5}

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(y[1:])  // Slice - from index 1 to end
	fmt.Println(y[1:3]) // Slice - from index 1 up to but not including

	for index, value := range y {
		fmt.Println(index, value)
	}

	y = append(y, 7)
	z := append(y[:1], y[3:]...) // Will modify y!

	fmt.Println(y)
	fmt.Println(z)

	c := make([]string, 10, 13)

	c[0] = "Person 1"
	c[1] = "Person 2"
	c[9] = "Person 10"
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	c = append(c, "Person 12")
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	c = append(c, "Person 13")
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	c = append(c, "Person 14")
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

	c = append(c, "Person 15") // Will double capacity because append exceeded previous capacity
	fmt.Println(c)
	fmt.Println(len(c))
	fmt.Println(cap(c))

}
