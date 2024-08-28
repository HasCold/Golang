package main

import (
	"fmt"
	"sync"
)

// Race Condition :- Suppose you are writing something on one memory space using by one thread and another thread at the same time comes in and say I want to write into the same memory space, obvisously we are get into the issues

// command := go run --race .  // -->> Check whether the race condition occur in your program

func main() {
	fmt.Println("----------------------Race Condition------------------------")

	// We will wait the goroutine by sync.WaitGroup method and access the WaitGroup by the pointer
	wg := &sync.WaitGroup{} // wg holds the reference of sync.WaitGroup struct or creates a pointer to the sync.WaitGroup struct

	// Basiclly the idea of using mutexes is that when we implement the transactions or Lock the memory address location
	// we are resolving the race condition by Mutex instead of WaitGroup
	mut := &sync.Mutex{} // mut holds the reference of sync.Mutex struct or creates a pointer to the sync.Mutex struct
	Rmut := &sync.RWMutex{}

	var score = []int{0}

	// The wg.Add(4) statement indicates that you're adding 4 to the sync.WaitGroup counter. This means you are expecting to wait for 4 separate operations (usually goroutines) to complete before the WaitGroup allows the program to proceed.
	wg.Add(4)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("One Goroutine")
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("Two Goroutine")
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("Third Goroutine")
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.Mutex, Rmut *sync.RWMutex) {

		// It is a good practice when you are working with transactions so lock the resources while writing some operations
		Rmut.Lock()
		fmt.Println("Score :-", score)
		Rmut.Unlock()

		wg.Done()
	}(wg, mut, Rmut)

	wg.Wait() // we are telling our main function that you have to wait until our goroutines will complete their job

	fmt.Println("After Completing the Goroutine operations :- ", score)
}
