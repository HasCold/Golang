package main

import "fmt"

const (
	first = iota // 0
	second
	third = 'a' // This will convert 'a' according to ASCII code 97 = a
	forth
)

func main() {
	a := 20 // It is a short hand method ; Compiler judge the data-type at a compile time.
	b := 30

	fmt.Printf("The value of first is %v, \n second is %v, \n third is %v and \n the forth is %v \n", first, second, third, forth)
	fmt.Println(test(a, b))

	fmt.Println("--------------------------IF_ELSE-------------------------------------")
	condition()

	fmt.Println("--------------------------SWITCH_CASE-------------------------------------")
	sw_case()
}

func test(num1 int, num2 int) (int, error) { // return type is (int, error)
	return num1 + num2, nil
}

func condition() {
	if rune('a') == third { // Basically rune is a method to convert any into ASCII code and it is also accept only one arg with a '' of a type int32
		fmt.Println("A is true")
	}

	if rune('d') == third {
		fmt.Println("A is true")
	} else {
		fmt.Println("A is false")
	}

	if rune('c') == third {
		fmt.Println("third is true")
	} else if 0 == first {
		fmt.Println("first is true")
	} else {
		fmt.Println("I m in else case")
	}
}

func sw_case() {
	var a interface{}
	a = "Some Data"

	// In Go, type assertions and type switches are used to determine the dynamic type of an interface value. Demonstrates a type switch, which allows you to handle different types that an interface can hold in a type-safe way.
	switch t := a.(type) { // it is a short hand method to declare the variable
	case int64:
		fmt.Println("Type is an integer : ", t)
	case float64:
		fmt.Println("Type is a float : ", t)
	case string:
		fmt.Println("Type is a string : ", t)
		// fallthrough :-  Not recommended to use fallthrough because it will forcefully true the next condition and seems to satisfy more condition will true
	case bool:
		fmt.Println("Type is a bool : ", t)
	default:
		fmt.Println("Type is unknown !")
	}

}
