package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt")


func main() {
	fmt.Println("Hello!")

	test := []byte("Hello World")

	encryptedPassword, err := bcrypt.GenerateFromPassword(test, bcrypt.MinCost)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(encryptedPassword)

	err	= bcrypt.CompareHashAndPassword(encryptedPassword, test)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Password is correct!")
	}

}
