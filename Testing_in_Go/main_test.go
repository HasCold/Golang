package main

import "testing"

func TestReturnChannelName(t *testing.T) { //  The function takes a single argument t of type *testing.T, which is used to log errors or other test-related messages during the execution of the test.
	actualOutput := ReturnChannelName()
	expectedOutput := "Hello, Golang !"
	if actualOutput != expectedOutput {
		t.Errorf("Expected String(%s) and their Type(%T) is not same as Actual String(%s) and their Type(%T)", actualOutput, actualOutput, expectedOutput, expectedOutput)
	}
}

// In Go, the syntax t *testing.T indicates that t is a pointer to a testing.T struct. Here's why it's written that way and what it means:
// 1. Efficiency:

// .. Passing by Reference: When you pass a pointer (*testing.T), you're passing the memory address of the original testing.T object, not a copy of it. This is more efficient because you're not duplicating the entire struct, especially if the struct is large.

// .. Direct Modification: By passing a pointer, the function can modify the original object that was passed in. This is important for updating the test state, like recording that a test has failed.
