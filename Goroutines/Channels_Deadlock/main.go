package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("--------------------- The Channels Deadlock Explanation -----------------------")

	// myCh := make(chan int) // Un-Buffered Channels
	myCh := make(chan int, 2) // Buffered Channels
	wg := &sync.WaitGroup{}

	// Channel says that I am only allowing you to pass me a value if somebody is listening
	// If somebody listening to channel then it will allow to pass something

	// myCh <- 5 // <-myCh = recieve-only channel and myCh<- = send-only channel

	wg.Add(2) // I will tell waitGroups that hey, I am adding two goroutines

	// Send-only Channel
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 100
		ch <- 55
		close(ch) // Data breaching, often referred to as a data breach, is an incident where unauthorized individuals gain access to confidential, sensitive, or protected information. This can occur when hackers bypass security systems, exploit vulnerabilities, or when data is accidentally exposed due to poor security practices

		wg.Done() // And after completing operation or task then Done
	}(myCh, wg)

	// go func(ch chan int, wg *sync.WaitGroup) {
	// 	// Here is the problem that if the channel close directly then it throws a value 0 so we haven't know about whther it's a signal or value so there is method to tacke the problem our channel give us two things :- value, isChannelOpen
	// 	close(ch)
	// 	wg.Done()
	// }(myCh, wg)

	// Recieve-only Channel
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// close(ch) // The channel cannot be closed becasue it is recieve-only channel or read only channel

		if val, isChannelOpen := <-ch; isChannelOpen {
			fmt.Println("The isChannelOpen output :- ", isChannelOpen)
			fmt.Println("The channel value output :- ", val)
		} else {
			fmt.Println("The isChannelOpen is closed :- ", isChannelOpen)
		}

		wg.Done()
	}(myCh, wg)

	wg.Wait() // I am telling main goroutine or function that wait for my guys they are just coming in.

}
