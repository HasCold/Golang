package main

import "fmt"

func main() {
	a := a()
	b := b()

	// There are two types of function
	// 1. Pass by value
	// 2. Pass by reference

	sum(a, b)
	sumByReference(&a, &b) // Ampersand (&) returns the memory address location of variable
}

func a() int {
	return 10
}

func b() int {
	return 17
}

func sum(a int, b int) {
	// Any modifications to a and b inside sum won't affect the variables a and b in the main function.
	fmt.Println("Pass by value sum :- ", a+b)
}

func sumByReference(a, b *int) { // "*" we will get the value reference from memory address and then perform some other operation in a separate memory location; When you pass a pointer to a variable (using *int), you're effectively passing a copy of the memory address where the variable is stored. This allows you to modify the original variable indirectly.
	fmt.Println("The reference sum of a and b is", *a+*b)
}
