package main

import (
	"fmt"
	"log"
	"os"
)

// File system in golang
// golang offers inbuilt library to create file along with the read and write files operation
// Lib Like :- io, os, bufio

// Interview Question :-
//  Suppose we have 1 gb file so how can we do the read and write operations
// --->>> We can do that from bufio	library  =  bufio.NewReader(os.Stdin)
// We use bufio so that our CPU performance has been increased

//
// Methods :-
// os.Create()
// os.ReadFile()
// ioutil.ReadFile()  --->>> I/O utility functions deprecated
// ioutil.WriteFile() --->>> I/O utility functions deprecated
// bufio.NewReader(os.Stdin)

func CreateFile() {
	fmt.Println("Writing to a file in Golang")

	file, err := os.Create("file.txt")
	if err != nil { // When we are dealing with the critical error that encounters in the program so we can log the error messages by the log.Fatal() method and then immediately terminates the program with non-zero exit status code
		log.Fatalf("Failed creating file: %v", err)
	}

	defer file.Close() // defer before the statement will delay in the execution or execute this statment after all the nearby function will complete thier execution

	len, err := file.WriteString("Hello Golang from Hasan")
	if err != nil {
		log.Fatalf("Failed writing to file: %v", err)
	}

	fmt.Printf("\n File Name: %v", file.Name())
	fmt.Printf("\n Length: %d bytes and their type(%T) \n", len, len)
}

func readFile() {
	fmt.Printf("\n \n Reading a File in Golang \n")
	fileName := "file.txt"

	data, err := os.ReadFile("file.txt")
	if err != nil {
		log.Panicf("Failed reading file: %v", err)
		return
	}

	fmt.Printf("\n File Name: %v", fileName)
	fmt.Printf("\n Size: %v bytes", len(data))
	fmt.Printf("\n Data: %v and their type(%T)", string(data), data)

	// If you have read the very large gb file then you must use the bufio library

}

func main() {
	// CreateFile()
	readFile()
}
