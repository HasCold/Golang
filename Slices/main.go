package main

import "fmt"

// Slices are similar like array but not array main reason are :-

// 1. Powerful, flexible than array and light weight data structure
// 2. Similar type elements with no length
// 3. Index value 0 to len(arr) - 1
// 4. Slices is a reference to an underlying array
// 5. Internally both are connnected
// 6. []data_type
// OR
// 6. []data_type{}
// 7. []data_type{value1, value2, value3,......value n}

// Components
// Pointers
// Length
// Capacity

func main() {
	// Creating an array
	arr := [5]string{"This", "is", "the", "Golang", "Course"}

	// fmt.Println("Array :- ", arr)
	// fmt.Println("Capacity Array :- ", cap(arr)) // 5

	// Creating a Slice
	mySlice := arr[1:4] // Containing an element from 1 to 3 index

	// // Print Slices
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) // 3 length
	fmt.Println(cap(mySlice)) // 4 capacity

	// Slices :-
	// slice1 := arr[2:5] // Slice Containing an element from 2 to 4 index
	// slice2 := arr[0:2] // Slice Containing an element from 0 to 1 index
	// slice3 := arr[:]   // Slice Containing all elements

	// Print the original array and slices
	// fmt.Println("Slice 1:", slice1)
	// fmt.Println("Length of Slice 1:", len(slice1))
	// fmt.Println("Capacity of Slice 1:", cap(slice1))
	// fmt.Println("Slice 2:", slice2)
	// fmt.Println("Slice 3:", slice3)

	fmt.Println("-------------------- Append the Slice index ------------------------------------")
	mySlice = append(mySlice, "test")  // First args must also be a slice
	mySlice = append(mySlice, "Hello") // If the capacity is 4 and if we are appending more elements into the slice then the capacity will be double means 4 * 2 = 8

	// Print Slices
	fmt.Println(mySlice)
	fmt.Println(len(mySlice)) // 5 length
	// capacity will increase by * 2
	fmt.Println(cap(mySlice)) // 8 capacity

	fmt.Println("-------------------- Diff between Array and Slice ------------------------------------")
	// In array we have to mention the length but in slices we haven't need to mention the length
	// Array = Fixed Size Length
	// Slice = Dyammic Size length or No need to mention the length when declaring the slice

	var my_Slice1 = []string{"Hello", "World", "Its", "Me"}
	fmt.Println("My SLice 1 :- ", my_Slice1)

	var my_Slice2 = []int{1, 2, 3, 4, 5}
	fmt.Println("My SLice 1 :- ", my_Slice2)

	//
	fmt.Println("------------------------ Slice From Make  ------------------------------------")
	var slice_make = make([]int, 4, 7)
	fmt.Printf("The Slice 1 :- %v,\n Length is %v, \n Capacity is %v \n", // len :- 4 and Cap :- 7
		slice_make, len(slice_make), cap(slice_make))

	var slice_make2 = make([]int, 5)
	fmt.Printf("The Slice 2 :- %v,\n Length is %v, \n Capacity is %v", // len :- 5 and Cap :- 5
		slice_make2, len(slice_make2), cap(slice_make2))

	//
	fmt.Println("------------------------ Slices is reference to an underlying array  ------------------------------------")
	arr2 := [6]int{1, 2, 3, 4, 5, 6}
	slice_1 := arr2[0:4] // Containing an elements from 0 to 3 index and slices are referenced to an underlying array

	fmt.Println("Before Changing the Slice the Array :- ", arr2)
	fmt.Println("Before Changing the Slice the Slice :- ", slice_1)

	slice_1[0] = 10
	slice_1[1] = 100
	slice_1[2] = 1000

	fmt.Println(".......")
	fmt.Println("After Changing the Slice the Array :- ", arr2)    // [10 100 1000 4 5 6]
	fmt.Println("After Changing the Slice the Slice :- ", slice_1) // [10 100 1000 4]

}
