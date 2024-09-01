package main

import (
	"errors"
	"fmt"
	"os"
)

func sum(a int, b int) (res int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("both a and b must be non-negative")
	}

	return (a + b), nil
}

type CustomError struct {
	data string
}

func (e *CustomError) Error() string { // pointer reciever
	return fmt.Sprintf("Error Occured due to %v", e.data)
}

func main() {
	r, err := sum(10, 10)
	if err != nil {
		var d CustomError
		d.data = err.Error()
		fmt.Println(d.Error())
		return
	}

	fmt.Printf("The result is %v and error is %v \n", r, err)

	file, err := os.Open("./abc.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(file)
}
