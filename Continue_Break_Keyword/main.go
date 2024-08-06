package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break // Exit the loop when i equals 5
		}
		fmt.Println(i)
	}
	fmt.Println("Loop ended")
}

func loop() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // Skip the rest of the loop body for even numbers
		}
		fmt.Println(i)
	}
	fmt.Println("Loop ended")
}
