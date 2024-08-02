package main

import (
	"fmt"
	"strings"
)

func main() {

	// Split String
	str1 := "Hello World, Hasan, Umer"

	// Will return strings array
	strArr := strings.Split(str1, ",")
	fmt.Println("After Split", strArr)

	// Split After
	strArr2 := strings.SplitAfter(str1, ",")
	fmt.Println("SplitAfter Array ", strArr2)

	// Split After N
	str3 := strings.SplitAfterN(str1, ",", 2)
	fmt.Println("SplitArray N :- ", str3[1])

	fmt.Println(strings.ToUpper(str1)) // Captialize the actual string
	fmt.Println(strings.ToLower(str1))
	fmt.Println(strings.ToTitle(str1)) // returns the copy of string

}
