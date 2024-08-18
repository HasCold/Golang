package main

import (
	"encoding/json"
	"fmt"
)

// JSON -->>> JavaScript Object Notation
// User readable format
// {
// 	"key": "value"
// }

// How to bind JSON data with object
// How to access keys
// ok syntax with JSON keys

type User struct {
	Name        string
	Designation string
	Phone       int64
}

func main() {
	// Binding JSON with Map method
	var wrapData map[string]interface{}

	// Binding JSON with Struct method
	// var wrapData User

	data := `{"name": "M.HasC Al", "designation": "programmer", "phone": 123456}`

	// json.Unmarshal(The data which you want to bind, The data which you want to bind with,just passed the reference of memory address)
	// Maps are referenced type, we have to pass their reference
	err := json.Unmarshal([]byte(data), &wrapData)

	if err != nil {
		fmt.Println("Error while fetching the JSON data :-", err)
	}
	// fmt.Println(wrapData["phone"])

	val, ok := wrapData["name"]
	if ok {
		fmt.Println("Name is :-", val)
	} else {
		fmt.Println("Not present key :-", ok)
	}
}

// NOTE :- you can use json.Marshal to get back data in JSON format => {"name": "example"}
