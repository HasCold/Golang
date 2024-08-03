package main

import "fmt"

func main() {
	// main() is also a goroutine so if interviewer asked how many goroutines will you able to determine so in this case main is also a goroutine.

	// Interview Questions :- Does init function contain any argument
	// Answer :- No, init function doesn't contain any argument or never return anything

	fmt.Println("The main function")
}

// Multiple Init Function
// Init function will execute first  before the main function
// e.g. :- Init function can be used to initialize the database connection and setting the route handler before our application run.
func init() {
	fmt.Println("Init First")
}

func init() {
	fmt.Println("Init First")
}
