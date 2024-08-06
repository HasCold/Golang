// In Go, make is a built-in function used to create and initialize slices, maps, and channels. It is essential for these three composite types because it allocates and initializes the internal data structures needed for them to function properly. Here’s a detailed look at the purpose and usage of make in Go:

// Purpose of make

// 1. Slices:
// make allocates an array and returns a slice that refers to that array.
// It initializes the slice with a specific length and capacity.

// 2. Maps:
// make allocates and initializes a hash map data structure.
// It prepares the map to store key-value pairs.

// 3. Channels:
// make creates and initializes a channel.
// It sets up the channel to handle communication between goroutines.

// Syntax
// The syntax for make varies slightly depending on the type you are creating.
// make(type, size, capacity) // For slices (capacity is optional)
// make(type, size)           // For maps and channels (size is optional for maps)

// ----------------------------------------- Slices ----------------------------------------
// 1. Slices
// package main

// import "fmt"

// func main() {
//     // Create a slice with length 5 and capacity 10
//     slice := make([]int, 5, 10)

//     fmt.Printf("Slice: %v, Length: %d, Capacity: %d\n", slice, len(slice), cap(slice))
// }
// In this example, make([]int, 5, 10) creates a slice of integers with a length of 5 and a capacity of 10.

// ----------------------------------------- Maps ----------------------------------------
// 2. Maps
// package main

// import "fmt"

// func main() {
//     // Create a map with an initial space for 10 key-value pairs
//     m := make(map[string]int, 10)

//     // Add some key-value pairs
//     m["foo"] = 1
//     m["bar"] = 2

//     fmt.Println("Map:", m)
// }
// In this example, make(map[string]int, 10) creates a map with an initial space for 10 key-value pairs. Note that the size parameter is optional and only hints the initial allocation size.

// ----------------------------------------- Channels ----------------------------------------
// 3. Channels
package main

import "fmt"

func main() {
	// Create a buffered channel with capacity 5
	ch := make(chan int, 5)

	// Send some values to the channel
	ch <- 1
	ch <- 2

	// Receive values from the channel
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

// In this example, make(chan int, 5) creates a buffered channel with a capacity of 5. This allows up to 5 values to be sent to the channel without blocking
