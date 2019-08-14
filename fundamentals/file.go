package main

// Packages
import (
	"fmt"
)

// Constants
const (
	ConstExample = "const before vars"
)

// Variables
var (
	ExportedVar    = 42
	nonExportedVar = "so say we all"
)

// Main Type(s)
type User struct {
	FirstName, LastName string
	Location            *UserLocation
}

type UserLocation struct {
	City    string
	Country string
}

// Functions
func NewUser(firstName, lastName string) *User {
	return &User{FirstName: firstName,
		LastName: lastName,
		Location: &UserLocation{
			City:    "Santa Monica",
			Country: "USA",
		},
	}
}

// Methods
func (u *User) Greeting() string {
	return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
}

func main() {

}
