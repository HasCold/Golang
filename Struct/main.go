package main

import "fmt"

// Structure / Struct
// It is a user-defined collection of data-type
// It is a collection of multiple data-type

type user struct {
	name    string      // By-default represent from the ""
	email   string      // By-default represent from the ""
	phone   int         // By-default represent from the 0
	address interface{} // By-default represent from the nil
}

// ----------------------------- Nested Struct --------------------------------
// Nested Struct
type user_address struct {
	street   string
	city     string
	division string
}

type user_details struct {
	name    string
	email   string
	address user_address
}

func main() {
	// var u user
	// fmt.Println(u)

	// One Way
	// u = user{"Hasan", "ha033@gmail.com", 12121121, "some address"}
	// fmt.Println(u)

	// // Two Way
	// u = user{
	// 	name:    "Hasan",
	// 	email:   "ha03330@gmail.com",
	// 	phone:   123123123,
	// 	address: "some interface",
	// }
	// fmt.Println(u)

	// ------------------------- Pointers With Struct -----------------------------
	u := &user{name: "Hasan", email: "umer@gmail.com"} // copy the memory address location
	fmt.Println("The Pointer name :- ", u.name)
	fmt.Println("The Pointer email :- ", (*u).email)

	// ----------------------------- Anonymous Struct --------------------------------
	// Anonymous structure
	// var name := struct{
	// 	// fields data type
	// }{fields_value}

	test := struct {
		int
		string
	}{123, "Hello"}

	fmt.Println(test)
}
