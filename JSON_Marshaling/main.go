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

	bytesData, err := json.Marshal(&wrapData) // return the response into the byte array and also return error if present

	if err != nil {
		fmt.Println("Error while Marshaling the data :- ", err)
	}
	fmt.Println(string(bytesData)) // we have to do type casting

	// json.Unmarshal(The data which you want to bind, The data which you want to bind with,just passed the reference of memory address)
	er := json.Unmarshal(bytesData, &wrapData)

	if er != nil {
		fmt.Println("Error while fetching the JSON data :-", er)
	}
	fmt.Println(wrapData)

}

// NOTE :- you can use json.Marshal to get back the data in JSON format => {"name": "example"}
