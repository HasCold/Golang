package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}
var mut = &sync.Mutex{}

func main() {
	// We can create a mutex for dealing the transaction to maintain the synchronization of the data and safely exchange data

	count := 0

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&count, wg, mut) // pass by reference or pointer not the pass by value or copy
	}

	wg.Wait()
	fmt.Println("Outside the worker", count)
}

func worker(count *int, wg *sync.WaitGroup, mut *sync.Mutex) {
	mut.Lock()

	*count++
	fmt.Println("Inside the worker", *count)

	mut.Unlock()
	defer wg.Done()
}
