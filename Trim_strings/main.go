package main

import (
	"fmt"
	"strings" // strings package or module
)

func main() {

	// It is a short hand method to declare the variables and compiler judges their data-type at a compile time.
	str1 := "@Some@ @String+@+@"
	fmt.Println(strings.Trim(str1, "@")) // Trim basically removes the character from start and at the end.

	fmt.Println(strings.TrimLeft(str1, "@"))
	fmt.Println(strings.TrimRight(str1, "@"))

	str2 := "       Some string    "
	fmt.Println(str2, "Before Print")
	fmt.Println(strings.TrimSpace(str2))

	// Trim suffix
	str3 := "Hello Friends"
	fmt.Println(strings.TrimSuffix(str3, "Friends"))

	// Trim prefix
	str4 := "Hello Friends"
	fmt.Println(strings.TrimPrefix(str4, "Hello"))
}
