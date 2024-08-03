package main

import "fmt"

func main() {
	// Variadic function : - No. of arguments you can passed and then access with a "..." spread operator along with mentioning the data-type
	// Variadic function returns an array
	variadicFunc(1, 2, 3, 4, 4, 5, 56, 6, 77, 7, 8, 8, 8, 8, 9, 6, 56, 5, 10, 5, 56, 5, 5, 56, 56)

	// Anonymous Function :- A function which hasn't any name
	// In JS, called as IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("Hello, World! Anonymous Function")
	}()

	// In JS, you can called as an arrow fucntion new EcmaScript Syntax
	val := func() int {
		return 11
	}
	fmt.Println("The value function :- ", val())

	// Variadic Function with Anonymous Function
	func(v ...string) {
		fmt.Println("Variadic Anonymous Fuction :- ", v)
	}("test1", "test2", "test3", "test4")
}

func variadicFunc(v ...int) {
	fmt.Println("Elements are", v)
}
