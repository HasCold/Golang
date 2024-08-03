package main

import "fmt"

// Closure Functions :- A closure is a function value that references variables from outside its body. The function may access and assign to the referenced varaibles; In this sense, the function is "bound" to the variables.

// When we use the closure function we need the anonymous function

// Closure functions are those function which has lexical scope and also returns another function with the use of its own lexical environment variable
func company() func() int {
	a := 1
	return func() int {
		a++
		return a
	}
}

func main() {
	val := company()
	fmt.Println("The val is :-", val()) // 2
	fmt.Println("The val is :-", val()) // 3
	fmt.Println("The val is :-", val()) // 4

	v := company()
	fmt.Println("The v is :- ", v()) // 2
}
