package main

import (
	"fmt"
	"strings"
)

// Strings are compared based on their ASCII code
// ASCII code is a mathematical representation of any character

func main() {
	str1 := "abc"
	str2 := "ab"

	fmt.Println(str1 == str2)

	// ASCII code
	// "98+12+97" > "1000"
	result := "ABC" > "xyz"
	fmt.Println("Result 1 :-", result)

	result2 := "XYZ" < "abc"
	fmt.Println("Result 2 :-", result2)

	fmt.Printf("The ASCII code of ABC and xyz are %v and %v", rune('A'), rune('x'))

	// Return 0 , if str1 == str2
	// Return 1, if str1 > str2
	// Return -1, if str1 < str2

	fmt.Println(strings.Compare("ABC", "ABC"))
	fmt.Println(strings.Compare("ABc", "ABx"))
	fmt.Println(strings.Compare("Ayz", "ABC"))
	fmt.Println(strings.Compare("BooKs", "BooKs"))
}
