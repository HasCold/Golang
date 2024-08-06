package main

import (
	"fmt"
	"time"
)

// Goroutines

// 1. Every process which run concurrently in golang called goroutines
// 2. Gouroutines is a light weighted thread
// 3. Creation cost of goroutines is very small as compared to thread
// 4. Every program had at leaast one single goroutine called main function
// 5. When main goroutine is terminated then all goroutine will terminated means all routine works under main

// prefix of your function invoking you can create goroutine
// go doSomething()

func main() {
	// go doSomeMagic() // we create a goroutine
	// doSomeMagic()

	// Anonymous Goroutine
	fmt.Println("Hello main")
	go func() {
		fmt.Println("Hello from anonymous goroutine")
	}()
	time.Sleep(1 * time.Second)
}

func doSomeMagic() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second) // Will run the goroutine at a 1 second interval
		// time.Sleep(1 * time.Minute)
		fmt.Println("In the loop", i)
	}
}
