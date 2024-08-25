package main

import "fmt"

// Difference between go build and go install

func main() {
	fmt.Println("Hello World, Hasan")
}

// Command :-
// echo $GOPATH
// go env  --->>> From this command we can find the GOPATH=C:\Users\Admin\go

//
// go build .  -->> Build all the files in the current directory which is present in a binary executable file
// check out the binary executable file -->> ./gog
// go build mainly compile the packages and dependency internally
//

//
// go install .   -->>> Install the module globally
// you don't need to specific path just put module name which is the go.mod -->>>  gog
// go install compile and also install dependency globally into the system
