package main

import (
	"log"
	"os"
	"fmt"
)

type Job struct {
	Command string
	Logger *log.Logger
}

type JobTwo struct {
	Command string
	*log.Logger
}

func main() {

	fmt.Println(log.Ltime) // 2
	fmt.Println(log.Ldate) // 1
	fmt.Println(log.Lmicroseconds) // 4
	fmt.Println(log.Llongfile) // 8
	fmt.Println(log.Lshortfile) // 16
	fmt.Println(log.LUTC) // 32
	fmt.Println(log.LstdFlags) // 3 - Ldate | Ltime


	job := &Job{"Stdout Logger Demo", log.New(os.Stdout, "Tranquility's Genesis: ", 1 | 16)}
	jobTwo := &JobTwo{"Stderr Logger Demo", log.New(os.Stderr, "Error: ", log.Ltime)}
	
	job.Logger.Print("Go to your destiny!")
	jobTwo.Print("File not found!")
}
