package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

// There is a pointer reciever and value reciver

func (p Person) Greet() {
	fmt.Println("Hello, my name is ", p.Name, "and I am", p.Age, "years old.")
}

func main() {
	// Third party call
	// Instance of the person struct
	p := Person{
		Name: "Alice",
		Age:  20,
	}

	t := reflect.TypeOf(p)  // Reflect the type
	v := reflect.ValueOf(p) // Reflect the values

	fmt.Printf("Type of t : %v and Value of v : %v", t, v)

}
