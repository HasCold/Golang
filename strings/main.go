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

	// fmt.Println(d[0]) // Returns the ASCII code of "S" string
	f := 'a' // rune is a data-type which returns a byte code or ASCII code typically required the value in a single quotes
	fmt.Println(f)

	// Interview Questions :- How does the back tick `` is differ from ""
	// In back ticks or template literals the escaping character like \n \t doesn't work ; they are considered to be a raw strings
}
