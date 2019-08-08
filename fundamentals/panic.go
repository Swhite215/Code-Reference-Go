package main

import "fmt"

func main() {
  fmt.Println("I'm First to Print")
  
  defer fmt.Println("I Will Print After Panic!")
  
  panic("I am the Panic Value!")
  
  fmt.Println("Final Print - I'm Skipped")
}
