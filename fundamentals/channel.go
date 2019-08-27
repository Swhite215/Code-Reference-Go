package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func sendChannel(c chan<- string) {
	for i := 0; ; i++ {
		c <- "message"
		// test := <- c // Causes an error because channel is send only
	}
}

func receiveChannel(c <-chan string) {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
		// c <- "test" // Causes an error becaue channel is receive only
	}
}

func main() {
	var c chan string = make(chan string)
	var d chan string = make(chan  string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	go sendChannel(d)
	go receiveChannel(d)

	var input string
	fmt.Scanln(&input)
}
