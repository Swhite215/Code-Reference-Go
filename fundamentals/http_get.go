package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	// Make a Basic Get Request with no Configuration
	resp, err := http.Get("https://httpbin.org/get")

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Defer the Close of the Response Body
	defer resp.Body.Close()

	// Response Fields
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Proto)
	fmt.Println(resp.ProtoMajor)
	fmt.Println(resp.ProtoMinor)
	fmt.Println(resp.Header)

	// Read the Response Body
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error:", err)
	}

	// Print the Response Body
	fmt.Printf("%s", body)
}
