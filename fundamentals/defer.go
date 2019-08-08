package main

// Function Using Defer to Ensure Files Are Closed
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

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
