package main
import (
	"flag"
	"math/rand"
	"fmt"
)

func main() {
 // Define Flags
 maxf := flag.Int("max", 6, "The Max Value")
 minf := flag.Int("min", -1000, "The Min Value")

 // Parse
 flag.Parse()

 fmt.Println(rand.Intn(*maxf))
 fmt.Println(rand.Intn(*minf))
}
