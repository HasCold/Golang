package main

import (
	_ "embed" // _ "embed": The underscore (_) before the package name indicates that you're importing the package for its side effects (in this case, enabling the use of the //go:embed directive) without directly referencing it in your code.
	"encoding/json"
	"fmt"
)

// go:embed -->> Deals with the static files and if there is a confidential things in the file so in this case I will convert them into binary information

//go:embed test_2.json
var d []byte // declarations following go:embed directives must be of type string, []byte or embed.FS

type Data struct {
	Name  string
	Phone string
}

func main() {
	// fmt.Println("The embedded data in byte array", string(d))  // This way to convert the byte array into the string
	fmt.Println("The embedded data in byte array", d)
	var marshalData Data

	json.Unmarshal(d, &marshalData)
	fmt.Println("The phone number is :- ", marshalData.Phone)

}

//
// Before build this file or convert the file into binary information
// 1. go mod init <GoEmbed>
// 2. go build .
