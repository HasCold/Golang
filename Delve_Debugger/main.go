package main

import "fmt"

// Delve  :- A debugger for the go programming language.

// Install the latest release:
// go install github.com/go-delve/delve/cmd/dlv@latest

// Command :-
// 1. dlv  ==  Binary command found

// Follow these steps :-
// - Ctrl + Shift + P
// - go install update
// - dlv@latest then click OK

// - Again Ctrl + Shift + P
// - Then Open Run and Debug

func main() {
	count := 0

	for {
		if count == 10 {
			break
		}
		count++
	}
	fmt.Println(count)
}
