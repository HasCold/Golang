package main

import "fmt"

// Topic :- Panic and Recover
// 1. System Panic
// 2. Custom Panic

// Panic create in the system due to some error happened so we can resolve it by the recover() method along the defer keyword function

func main() {
	x := 0
	y := 20
	printAllOperations(x, y)
}

// func printOperation(y, x int) { // panic: runtime error: integer divide by zero ; goroutine 1 [running]:
// 	panic("Custom Panic")
// 	fmt.Println(y / x)
// }

func printAllOperations(x, y int) {
	defer func() { // when used the defer keyword so the function execution will slow or delay in time and invoked until the nearby function is called
		// defer function to escape the panic when y/x

		// Recover method is used when our program is panic due to some error
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in printAllOperations error is : %v \n", r)
			fmt.Println("Proceeding to alternative flow skipping division.")
			printOperationSkipDivide(x, y)
		}
	}()

	sum, divide, multiply := x+y, y/x, x*y
	fmt.Printf("The sum = %v, divide = %v and multiply = %v \n", sum, divide, multiply)
}

// if Block :-
// 1. recover: This function returns nil if the goroutine is not panicking. Otherwise, it returns the argument passed to panic.
// 2. if r := recover(); r != nil: This line checks if there is a panic occurring. If there is a panic (i.e., recover() does not return nil), the code inside the if block is executed.

//
func printOperationSkipDivide(x, y int) {
	sum, mult := x+y, x*y
	fmt.Printf("The sum = %v and multiply = %v \n", sum, mult)
}

// Context of recover in Go :-
// . panic: When a Go program encounters a situation it cannot handle, it can call panic, which stops the normal execution of the program.
// . recover: This is used to regain control of a panicking goroutine. When called inside a deferred function, recover can stop the panic and return the value passed to panic.
