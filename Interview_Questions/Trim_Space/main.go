package main

import (
	"bufio"
	"fmt"
	"os"
)

func TrimSpace(input string) string {

	start := 0
	end := len(input) - 1

	// ASCII Values or Code -->> The mathematical representation of any character
	for start <= end && input[start] == ' ' {
		start++
	}

	for end >= start && input[end] == ' ' {
		end--
	}

	return input[start : end+1] //  input[0 : 3]  means return the value from 0 to 2 index
}

func main() {
	fmt.Println("Enter the sentences with some spaces before and after !")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	trimInput := TrimSpace(input)
	fmt.Println("Trim Input :- ", trimInput)
}
