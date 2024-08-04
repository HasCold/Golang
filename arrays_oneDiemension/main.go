package main

import (
	"fmt"
	"reflect"
)

// arrays are similar like java, node, c++
// Arrays contains similar type of data
// It has fixed length
// Due to the fixed length they are not much popular in golang // slices
// index operator = []

func main() {
	// In Go, when you declare an array, its size and type are fixed. However, the array elements need to be initialized, either explicitly or implicitly, with zero values for the type (e.g., 0 for integers, "" for strings).

	// Types of arrays :- One diemension and Multi diemension arrays

	// Type 1
	// arr1 := [3]int{1,2,3}   // Declaration and initialization of an integer array with specific values
	var arr1 [3]int

	// Type 2 : Declaration and implicit initialization with zero values
	arr2 := [3]string{} // Array of strings with zero values (empty strings)

	arr1[0] = 1
	arr1[1] = 2
	arr2[0] = "1"
	arr2[1] = "2"
	arr2[2] = "3"

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i], reflect.TypeOf(arr2[i]))
	}

	// Interview level question
	t1 := [3]int{9, 7, 6}
	t2 := [...]int{9, 7, 7, 6, 5, 54, 4} // [...] it defines the fixed length of an array by the spread operator when we initilize into the index operator []

	fmt.Println("Length of an array 1 is:- ", len(t1))
	fmt.Println("Length of an array 2 is:- ", len(t2))

	// Comparison in Arrays
	q1 := [3]int{9, 6, 7}
	q2 := [...]int{9, 6, 7}
	q3 := [3]int{9, 5, 3}

	fmt.Println(q1 == q2)
	fmt.Println(q2 == q3)
	fmt.Println(q1 == q3)

	// --------------------------------- Array in Depth ------------------------------
	// How to copy array using refernces and without reference
	// Pass array to function

	array1 := [5]int{0, 10, 2, 3, 4}

	// Without reference :- Means both variable has their own memory address location
	// array2 := array1
	// fmt.Println("Simple array value copy :-", array2)

	// array2[0] = 12
	// fmt.Println("After changing the value of array  :-", array2)
	// fmt.Println("After changing the value of array  :-", array1)

	// With reference :- Means both variable has access of same memory address location
	// So it will direclty impact on our actual variable
	array2 := &array1 // Both have access the same memory address location
	array2[1] = 100

	fmt.Println("After changing the value of array  :-", *array2)
	fmt.Println("After changing the value of array  :-", array1)

	testArrayPassing(array1)
}

func testArrayPassing(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(i)
	}
}
