package main

import (
	"fmt"
	"sync"
	"time"
)

// Methods in Golang :- The methods that can wait the goroutines and are used to print the goroutines.
// Topic :- waitGroup

// ---------------- Methods ---------------------
// 1. waitGroup
// 2. time.Sleep(time.Second * 5) --->>>  This is not the recommended method
// 3. Synchronize Channel
// Basically compiler will not wait for the goroutines

var wg sync.WaitGroup

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Wait Groups Worker %v starting \n ", id)
	time.Sleep(time.Second * 3)
	fmt.Printf("Wait Groups Worker %v finished \n", id)
	fmt.Println("--------------------------------------------")
	wg.Done()
}

func worker2(id int, check chan bool) {
	fmt.Printf("Channel Worker %v starting \n ", id)
	time.Sleep(time.Second * 3)
	fmt.Printf("Channel Worker %v finished \n", id)
	check <- true // send-only channel
}

func main() {
	chan1 := make(chan bool)

	for i := 0; i <= 5; i++ {
		wg.Add(1) // we are telling that we have one goroutine

		go func(channel chan bool) {
			// defer wg.Done() // defer keyword will run this line of code at the end of a execution or make some delay
			worker(i, &wg)
			worker2(i, channel)
		}(chan1)
		fmt.Printf("The channel value is %v \n", <-chan1) // This is the recieve-only channel
	}
	// time.Sleep(5 * time.Second)
	wg.Wait() // Basically we have to wait for the goroutines that have added for e.g. we added 1 goroutine
}
