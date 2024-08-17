package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan int)

	go funcTwo(chan2)
	go funcOne(chan1)

	// Select statement basically works on the channel which comes first and execute their corresponding case
	select {
	case val1 := <-chan1: // Recieve-only channel;
		fmt.Println("The value 1 is :-", val1)

	case val2 := <-chan2: // Recieve-only channel;
		fmt.Println("The value 2 is :-", val2)

	default:
		fmt.Println("This is the default value")
	}

	// select {
	// default:
	// 	fmt.Println("This is the default case")
	// }
}

func funcOne(ch1 chan string) {
	ch1 <- "Hello Golang" // send-only channel
	time.Sleep(time.Second * 5)
}

func funcTwo(ch2 chan int) {
	ch2 <- 42                   // send-only channel
	time.Sleep(time.Second * 2) // This channel case will return the int first
}
