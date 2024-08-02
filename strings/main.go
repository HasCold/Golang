package main

import "fmt"

// type 1 :- ""
// type 2 :- ``  // Template-literals

func main() {
	// Strings are read only and are immutable

	// Short-hand method
	a := "golang"

	// var keyword
	var b string
	b = "I like coding"

	fmt.Println("Concatenate two string :- ", a+" "+b)

	// Escaping character  :-  \n, \t

	d := "Some String"
	fmt.Println("Some \t", d) //  \t gives the tab/space

	// raw literals
	c := `go golang asasjkdasjbaj \n \t `
	fmt.Println("The value of c :-", c)

	// We can access the strings by the index
	// fmt.Println(d[0]) // Returns the ASCII code of "S" string

	// f := 'a' // rune is a data-type which returns a byte code or ASCII code typically required the value in a single quotes
	// fmt.Println(f)

	// Interview Questions :- How does the back tick `` is differ from ""
	// In back ticks or template literals the escaping character like \n \t doesn't work ; they are considered to be a raw strings

	// ------------------------- Range Loop ------------------------------

	// var testStr string = "test method"
	testStr := "test method"
	for index, char := range testStr { //  index ;  char = bytes / ASCII code
		fmt.Println("\n index", index, "\n bytes", char, "\n characters", string(char)) // Type assertion string(char)
	}

	fmt.Println([]byte(testStr)) // type assertion show the strings into bytes / ASCII code
}

// What is Type Assertion?
// Type assertion is a way to explicitly specify the type of a value, telling the compiler to treat the value as a particular type. It doesn't change the runtime type or value; it just gives the compiler more information about what to expect.

// When is Type Assertion Used?
// Type assertion is useful when the compiler isn't able to infer the type correctly but you, as the developer, know the type. This often happens when working with any type or when using third-party libraries where type information might be lost or not detailed enough.

// ----------------------------------------------------

// Type Assertion and Type Casting

// Purpose:
// Type assertion is used to tell the compiler about the type of a value when you are sure of its type.
// Type casting is used to convert a value from one type to another.

// Runtime Behavior:
// Type assertion does not change the actual type or value at runtime; it only affects how the compiler treats the value.
// Type casting changes the actual type and possibly the value representation at runtime.
