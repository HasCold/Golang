package main

import "fmt"

// Pointers :- Pointers are basically used to store references of one variable in to other variable. Pointers are in hexadecimal format which has a prefix of "0x". Basically every variable has stored into our memory address and thier location looks something like this 0x12332132 so in this case we use pointers in our programs

// In Go, pointers are variables that store the memory address of another variable. They allow you to reference and manipulate the values stored in different locations in memory.

func main() {
	// & (Ampersand)  is used for the comsumption of memory address or to access the memory address
	// * is used to get the value from memory address

	var a *int
	var b *int

	var str1 *string // * to get the value from memory address location
	var str2 *string // * to get the value from memory address location

	a = ptr1(10) // 0xc000100038  memory address location
	b = ptr1(20) // 0xc000100070  memory address location

	c := *a + *b
	fmt.Println(c)

	str1 = ptr("10")
	str2 = ptr("20")

	str := *str1 + *str2 // In this way we can concatenate the two strings
	fmt.Println(str)
}

// --------------------- Go Generics -------------------------------
// Go generics allow you to write functions and data structures that can operate on any data type, providing type safety and code reuse.
func ptr1[T any](v T) *T { // Generics
	return &v
}

func ptr(v string) *string {
	// fmt.Println(&v)  // -->> print the memory address location
	return &v // consume the memory address and  return the memory address location
}

// The original use of pointers was unnecessary for this specific operation since you were simply adding two integers. Pointers are useful when you want to work with references to data (like when you want to modify the original data or avoid copying large amounts of data). In this case, since you're just returning and adding the values, pointers aren't needed, and the result remains the same when you remove them.
