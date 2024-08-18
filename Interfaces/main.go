package main

import "fmt"

// Polymorphism and Data Abstraction
// Polymorphism :- Poly means many and morphism means the type ; many-types
// Interfaces :- Interfaces are just like templates in which we declaring the functions without the implementation of their logic
// Interfaces :- Collections of methods and signatures

type values struct {
	first  int
	second int
}

// Data Abstraction means we provide you the interface and you don't worry about functionality
type mathTest interface {
	add(a, b int) int
	mul(a, b int) int
}

func (v values) add(a, b int) int { // Value Recievers
	return a + b + v.first + v.second
}

func (v values) mul(a, b int) int { // Value Recievers
	return a * b
}

func main() {
	var te mathTest = values{1, 2}
	fmt.Println(te.add(1, 2))
	fmt.Println(te.mul(1, 2))

	var test interface{} // Interfaces is just like the bucket ; means the latest value of test will be achieved
	test = "some string"
	test = 1
	fmt.Println("Here is the value of test :- ", test)

	// Type casting means convert one data-type to other data-type
	val, ok := test.(string) // This will check the condition if the test has string data-type then ok will be true
	if ok {
		fmt.Println(val)
	}
}
