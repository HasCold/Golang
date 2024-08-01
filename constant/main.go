package main

import (
	"fmt"
	"reflect"
)

func main() {
	const name = "Hasan"

	const (
		age   = 25
		grade = iota // iota contains the value related to the indexing like at 1 index iota value will be 1
		rollno
		class
	)

	fmt.Printf("The age is %v, \n grade is %v and \n class roll no is %v, %v",
		age, grade, class, rollno)

	fmt.Println("----------------------------------------------------------")
	fmt.Println("The type of name is ", reflect.TypeOf(name))

}
