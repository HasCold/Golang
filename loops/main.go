package main

import "fmt"

func main() {
	// for initialization; condition; post increment and decrement {
	// }

	// for i := 0; i < 4; i++ {
	// 	fmt.Println(i)
	// }

	// Infinite Loop
	// for {
	// 	fmt.Printf("Hello, World!")
	// }

	// While loop in Golang
	// j := 0
	// for j < 10 {
	// 	fmt.Println(j)
	// 	j++
	// }

	//  Range loop basically run over the array or slice
	// rvariable := []string{"0", "1", "2", "3"}

	// for i, val := range rvariable { //  i is index and val is index value
	// 	fmt.Printf("Index: %v and Value: %v \n", i, val)
	// }

	// Range loop supports over String and Print the value of string in ASCII (mathematical representation of any character) code
	// for index, value := range "Go, Golang World" {
	// 	fmt.Printf("The index number :- %v and their value is :- %v \n", index, value)
	// }

	// loop on Channels
	chanl := make(chan int) // Declare the channels

	go func() { // Use of Goroutine
		chanl <- 10
		chanl <- 100
		chanl <- 1000
		chanl <- 10000
		close(chanl)
	}()

	for i := range chanl {
		fmt.Println(i)
	}

	fmt.Println("------------------------------------------------------------------------------")

	// Loops range On Object

	mmap := map[int]string{
		0: "a",
		1: "b",
		2: "c",
	}

	for key, val := range mmap {
		fmt.Printf("Key :- %v and Value :- %v \n", key, val)
	}
}
