package main // represents the main package

import (
	"fmt" // module
	"reflect"
)

func main() { // main function is the entry point in our golang program. The code execution starts from main function
	fmt.Println("Hello Golang") // fmt.Println  --->> Print line

	var a int    // data-type declared
	var b = "20" // Compiler will judge data-type at a compile time
	c := 30      // := is a short hand method to declare a variable

	fmt.Printf("a is a typeof %v \n and b is a typeof %v \n and c is a typeof %v \n", // Printf used for formatting
		reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c))

}
