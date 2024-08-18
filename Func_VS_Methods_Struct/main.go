package main

import "fmt"

//  Important to Understand the Methods in Golang Deeply :- https://golangdocs.com/methods-in-golang

// Receiver types
// There are two types of receivers that are available in Go. The value receivers and the pointer receivers. Here is an example showing both of them.

type data int
type fl float32

func (v1 data) div(v2 data, v3 fl) fl { // Value Recievers
	return fl(v1/v2) + v3
}

// Struct is basically a user-defined collection of data-type
type user struct {
	name  string
	email string
	num   int
}

func (u *user) correctEmail(newEmail, newName string) { // Pointer Recievers
	u.email = newEmail
	u.name = newName
}

func main() {
	// Methods In Golang
	// func (receiver_name Type) method_name(parameter_list Type) (return_Type){

	// }
	// Methods have reciever and functions haven't anything but we can return the function
	// Inside the div method, v1 (the receiver) is value1, and v2, v3 (parameter list) are the argument value2, value3.

	// Creating instances of data and fl types ; Means type conversion
	// This is essentially calling the type conversion function data() and fl() to create instances of type data and fl from the integer literals 23, 20 and 12.6.
	value1 := data(10)
	value2 := data(2)
	value3 := fl(12.6)

	// Invoking the div method
	// Methods in Go have the ability to use the receiver instance. The receiver instance is the instance of the type on which the method is called. It allows the method to access and modify the fields of that instance.
	res := value1.div(value2, value3)
	fmt.Println("Final Result :- ", res)

	//  ------------------------------ Struct Data-Type ----------------------------------
	result := user{
		name:  "Hasan",
		email: "ha033@gmail.com",
		// num: 033022122323,
	}

	fmt.Println("User's name :- ", result.name)
	fmt.Println("User's email :-", result.email)

	// ------------------------------ Pointer with Methods --------------------------------

	// Creating a pointer
	p := &result // copy the memory address location

	// Calling the Show method
	p.correctEmail("ha033223732@gmail.com", "M.Hasan Ali") // Here we use the pointer with method correctEmail
	fmt.Println("user's name After:-", result.name)
	fmt.Println("user's email After :-", result.email)

	// ------------------------------ Return the function --------------------------------
	// Closure Function :- Those function with a lexical scope and also returns another function with the use of its own lexical environment varioble
	resp := simple()
	fmt.Println("The Simple function response is :- ", resp(1, 2))

}

func simple() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

// ---------------------------------------  Method Definition  --------------------------------

// func (v1 data) div(v2 data) data {
//     return v1 / v2
// }
// This defines a method div on the type data. The method signature has a few key components:

// 1. func (v1 data): This part defines div as a method with a receiver of type data. The receiver v1 is like a special parameter that allows the method to access the instance of data on which it's called.
// 2. div(v2 data): This part defines the method name div and the parameter list (v2 data), where v2 is also of type data.
// 3. data: This indicates the return type of the method, which is also data.
