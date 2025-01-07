package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	// Create a pointer to a Person struct with zero values
	p := new(Person)

	// The memory allocated is initialized to zero values
	fmt.Println(*p) // Output: { 0}
}

func Make() {
	// Using make to initialize a slice
	s := make([]int, 5, 10)

	// Modify the slice
	s[0] = 10

	// Print the slice
	fmt.Println(s) // Output: [10 0 0 0 0]
}

// In Go, both `new` and `make` are used to allocate memory, but they serve different purposes and behave in distinct ways. Here's a breakdown of their differences:

// ### Key Differences:

// | Aspect               | `new`                                             | `make`                                              |
// |----------------------|---------------------------------------------------|-----------------------------------------------------|
// | **Type**             | Allocates memory for any type, returns a pointer  | Initializes slices, maps, and channels, returns the initialized value |
// | **Return Value**     | Pointer to the allocated memory (e.g., `*T`)      | Initialized value (e.g., `[]T`, `map[K]V`, `chan T`)|
// | **Use Cases**        | Used for basic types, structs, arrays             | Used for slices, maps, and channels                |
// | **Initialization**   | Memory is initialized to zero values              | Internal structures are initialized (e.g., length and capacity for slices) |
// | **Example Use Case** | Creating a pointer to a struct                    | Creating a slice with specific length and capacity |

// ### Summary:

// - **`new`**: Used for creating a pointer to a new type (zero-initialized).
// - **`make`**: Used to initialize slices, maps, and channels with internal state (e.g., length, capacity) for immediate use.
