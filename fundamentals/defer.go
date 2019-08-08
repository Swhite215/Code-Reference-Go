package main

func main() {

  // Defer Leads to Stack LIFO
  fmt.Println("Hello")
  for i := 1; i <= 3; i++ {
      defer fmt.Println(i)
  }
  fmt.Println("World")
  
  // Prints: Hello -> World -> 3 -> 2 -> 1
  
  // Defer Can Alter Named Return Values
  func c() (i int) {
    defer func() { i++ }()
    return 1
  }
  
  // Returns: 2

}
