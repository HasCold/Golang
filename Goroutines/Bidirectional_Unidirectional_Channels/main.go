package main

import "fmt"

func main() {
	// ------------------------------------- Bi-Directional Channel -------------------------------------------
	// Creating a Bi-directional channel == We can put the value in channel or recieved a value from the channel
	chan1 := make(chan string)
	chan2 := make(chan string)

	go sending(chan1)           // creating goroutine
	valueFromChannel := <-chan1 // Opening the channel
	fmt.Println("Value from the channel :- ", valueFromChannel)

	go recieving(chan2)
	chan2 <- valueFromChannel

	// Sign Understanding
	// chan3 := make(chan<- string) // This is the send only channel in which we through the value to the channel
	// chan4 := make(<-chan string) // This is the recieved only channel in which we can get the value from the channel

	//
	// ------------------------------------- Uni-Directional Channel -------------------------------------------
	// chan4<-  are send-only channels
	chan4 := make(chan string) // Creating the channel
	go convert(chan4)          // Creating the goroutine and opening the channel
	fmt.Println("Send only Channel -------------", chan4)
}

func convert(s chan<- string) { // Send only channel
	s <- "Some string"
}

func sending(s chan string) {
	s <- "Go Golang !"
}

func recieving(s chan string) {
	fmt.Println("value comming from channel 2 _______________ ", <-s)
}
