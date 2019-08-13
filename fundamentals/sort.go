package main

import (
	"fmt"
	"sort")

type person struct {
	first string
	age int
}

type ByAge []person

func (a ByAge) Len() int { return len(a)}
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i]}
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age}

func main() {
	s := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(s)
	fmt.Println(s)

	p1 := person{"James", 32}
	p2 := person{"Spencer", 26}

	people := []person{p1, p2}

	fmt.Println(people)
	
	sort.Sort(ByAge(people))

	fmt.Println(people)
}
