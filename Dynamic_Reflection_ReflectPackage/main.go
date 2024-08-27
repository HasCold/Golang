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

	fmt.Printf("Type of t : %v and Value of v : %v \n", t, v)

	// Iterate through the struct fields using reflection
	fmt.Println("Field of person struct")
	fmt.Println("The value T :- ", t)
	fmt.Println("The value V :- ", v)

	fmt.Println("No of person struct field", t.NumField())

	for i := 0; i < t.NumField(); i++ {
		fmt.Println("The value t.Field :- ", t.Field(i))
		field := t.Field(i)
		fmt.Println("The value v.Field :- ", v.Field(i))
		fieldValue := v.Field(i)
		fmt.Printf("\n%s: %s \n\n", field.Name, fieldValue.Interface())
	}

	fmt.Println("-----------------------------------------------------------------------")
	// Call a method dynamically based on user input
	methodName := "Greet"
	// v is a person struct which we get
	method := v.MethodByName(methodName)

	// If the method is valid or exist then it will be called
	if method.IsValid() {
		fmt.Println("Calling method", methodName)
		method.Call(nil)
	} else {
		fmt.Println("Method", methodName, "Not Found.")
	}
}
