package main

import "fmt"

// Testing in Golang
// Manual Testing
// Automation Testing
// Unit Test Cases
// Integration Testing
// Load Testing

// How to do Testing in Golang ?
// Create a file for testing in golang
// <any-file-name>_test.go
// func TestReturnChannelName(){}

// By convention, test functions in Go are named starting with Test followed by the name of the function being tested
//

func ReturnChannelName() string {
	return "Hello, Golang !"
}

func main() {
	channelName := ReturnChannelName()
	fmt.Println(channelName)
}
