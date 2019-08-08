package main

import "fmt"

func main() {
  
  fmt.Println("I'm First to Print")

  defer func() {
    r := recover()
	fmt.Println("Program Recovers from Panic Here\n", r)
  }()
  
  panic("This is the Panic Value!")


  fmt.Println("Final Print - I'm Skipped")
}
