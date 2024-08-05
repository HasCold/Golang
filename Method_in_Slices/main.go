package main

import (
	"bytes"
	"fmt"
)

// Composite Literal
// Functions Like :- Sorted, Trimming, isorted. slice of bytes

func main() {
	// Slice with composite literals
	// sl := []int{1, 2, 3, 4, 5}
	// fmt.Println(sl)

	// Sorting ints
	// sl := []int{2, -32, -10, 22, 100, 12, 24, -3, 1, 56, 7}
	// sort.Ints(sl)
	// fmt.Println("Sorted Slice :- ", sl)
	// fmt.Println("Check Ints are sorted or not :- ", sort.IntsAreSorted(sl)) // Return either true or false

	//
	// The Slice of Bytes
	sl := []byte{'!', '#', 'G', '$', 'o', 'u', '&', '#', '!'}
	for i := 0; i < len(sl); i++ {
		fmt.Printf("The sl byte array will show the strings in a single quotation mark according to the ASCII code %v \n", sl[i])
	}

	//
	// Triming in Bytes
	resp := bytes.Trim(sl, "!#") // Trim these ASCII code (mathematical representation of characters)
	fmt.Println("The bytes resp :- ", resp)
}
