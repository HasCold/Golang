// Create our own custom packages

package custom1

import "fmt"

// Create a global variable and global variable cann't be created by short-hand method like this :- c :=
var Val int
var val int

func PrintValue(s string) {
	fmt.Println("The Val :-", Val)
	val = Val
	fmt.Println("Now the in-reserved val has same value :-", val)

	printVal(s)
}

//
// NOTE :- Functions or variables created with small letters can't access outside the package
// Example :-

func printVal(data string) {
	fmt.Println("In reserved function")
	fmt.Println("Can't accessible or expose outside the package :- ", data)
}
