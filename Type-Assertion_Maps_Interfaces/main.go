package main

import "fmt"

// Interface can hold any value regardless of their data-type
// Interface are collection of methods and signature ; Interface are just like your templates
type sumInterface interface {
	sum() int
}

type values struct {
	val1 int
	val2 int
}

// Method
// (reciever_name reciever_type)
// pointer reciever or value reciever
func (v values) sum() int { // value reciever
	return v.val1 + v.val2
}

// Mapping :- In mapping there is a hashing data structure is used ; In mapping, there is no any sequential order is followed
// This will called as a Hash Map
// {
// 	"key": "value"
// }

func main() {
	var a values
	a.val1 = 10
	a.val2 = 20
	fmt.Println("The sum of a.val1 and a.val2 is", a.sum())

	// Another tricky method
	var tricky sumInterface
	// Creates a pointer struct which is better and efficient way
	tricky = &values{val1: 100, val2: 200}
	fmt.Println("The tricky answer is that :- ", tricky.sum())

	//
	// ---------------------------------- Empty Interface -------------------------------------
	var test interface{}

	test = "Hello, Golang"
	test = &values{val1: 11, val2: 22} // Create a pointer to the value struct
	print(test)

	test = 10
	print(test)

	//
	// ---------------------------------- Map -------------------------------------
	testMap := make(map[int]int) // [key]value  // make method is a built-in data-type in golang is used to create and initialize the different data-type if we don't initialize the map from make keyword then the system will panic because make will allocate some memory address to the map initially

	testMap[0] = 10
	for k, v := range testMap {
		fmt.Printf("The key : %v and value is : %v", k, v)
	}
}

func print(t interface{}) {

	// Type assert to the specific type
	if v, ok := t.(*values); ok {
		fmt.Printf("The dereference pointer interface : %v and their status %v \n", *v, ok) // Dereference the pointer to remove '&'
	} else {
		fmt.Printf("The empty interface hold value : %v and their status %v \n", t, ok) // For non-pointer types
	}

}

// --------------------------------- Type Assertion ------------------------------------------

// ### Type Assertion

// ```go
// if v, ok := t.(*values); ok {
//     // v is now of type *values
// }
// ```

// - **Type Assertion**: In Go, a type assertion is used to check whether an interface value holds a specific type. The syntax `t.(*values)` checks if the `interface{}` variable `t` holds a value of type `*values` (a pointer to a `values` struct).

//   - **`v, ok := t.(*values)`**:
//     - `t.(*values)`: This tries to assert that `t` is of type `*values`.
//     - `v`: If the assertion succeeds, `v` is assigned the value stored in `t`, but with the type `*values`.
//     - `ok`: This is a boolean that is `true` if the assertion was successful (i.e., if `t` actually holds a `*values`), and `false` otherwise.

// - **Why is this necessary?**: Since `t` is an `interface{}`, it can hold any type of value. Before we can work with `t` as a `*values`, we need to ensure that it is indeed a `*values`. The type assertion allows us to safely do this check.

// ### Dereferencing

// ```go
// fmt.Printf("The dereference pointer interface : %v \n", *v)
// ```

// - **Dereferencing**: In Go, a pointer is a variable that holds the memory address of another variable. Dereferencing a pointer means accessing the value stored at the memory address the pointer is pointing to.

//   - `*v`: If `v` is a pointer (`*values`), dereferencing it with `*v` gives us the value of the struct that `v` points to.

// - **Example**: If `t` holds `&values{val1: 11, val2: 22}`, the type assertion `v := t.(*values)` makes `v` a `*values`. Dereferencing it (`*v`) yields the actual struct value `{val1: 11, val2: 22}`.

// ### `else` Block

// ```go
// } else {
//     fmt.Printf("The empty interface hold value : %v \n", t)
// }
// ```

// - If the type assertion fails (i.e., `t` does not hold a `*values`), the program enters the `else` block and prints `t` directly. This handles cases where `t` might hold a different type, such as a string, an integer, or another type.

// ### Summary

// - **Type Assertion**: Safely checks if `t` is a `*values` and retrieves the value if true.
// - **Dereferencing**: Accesses the actual struct value pointed to by `v` (`*v`).
// - **`else` Block**: Prints the value directly if it's not a pointer to a `values` struct.

// This approach ensures that if `t` holds a pointer to a struct, the actual struct value is printed, removing the `&` symbol, and if not, the value is printed as-is.
