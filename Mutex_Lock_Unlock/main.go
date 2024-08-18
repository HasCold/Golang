package main

import (
	"fmt"
	"sync"
	"time"
)

// Topic :- Mutex Lock
// Example :- Suppose we do a transaction of deposit 1000 in the bank and we cash out 1000 from the bank at a same time so this asynchoronization will remove from the Mutex Lock and Mutex Lock will convert this into Synchoronization also called as Mutual Exclusion.

// SafeCounter is safe to use concurrently
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int // var myMap map[keyType] valueType
}

// func (receiver_name Type) method_name(parameter_list Type) (return_Type){
func (c *SafeCounter) Inc(key string) { // Pointer recievers

	// Whenever this increment function will call this mutex will lock the transactions
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // defer keyword execute with some delay or execute at the end when all the nearby function will already invoked; Here when we return the int, defer keyword will execute along.
	return c.v[key]
}

func main() {
	// Now we have matched the struct data-type with method recievers data-type so we can call the function directly
	c := SafeCounter{v: make(map[string]int)}

	for i := 0; i <= 500; i++ {
		go c.Inc("Some key") // we have used the goroutines means I have created the 500 threads that will run independently
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("Some key"))
}
