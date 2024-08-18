package main

import "fmt"

// In Go, map is a built-in data structure that maps keys to values. It's similar to dictionaries in Python or hash maps in other languages. Here are some key points about map in Go:

// Declaration: Maps are declared using the map keyword followed by the key and value types enclosed in square brackets ([]).
// var myMap map[keyType]valueType

// For Example :-
// var scores map[string]int

//
// Initialization: Maps need to be initialized before use. This is typically done using the make function:
// scores = make(map[string]int)
// OR
// scores := make(map[string]int)

//
// Adding and Accessing Elements:

// To add or update a value, use the key within square brackets:
// scores["Alice"] = 92

// To retrieve a value, use the key:
// fmt.Println(scores["Alice"]) // prints 92

//
//
// Checking Existence: You can check if a key exists in the map using multiple assignment:
// score, exists := scores["Alice"]
// if exists {
//     fmt.Println("Alice's score is", score)
// } else {
//     fmt.Println("Alice's score is not recorded")
// }

// Deleting Elements: Use the delete function to remove a key from the map:
// delete(scores, "Alice")

// ------------------------------ Iterating Over a Map (Map VS Array)-----------------------------
// In map there is a hashing data-structure is used
// In map there is no any sequential order is followed where as in array there is indexing data structure

// Iterating Over a Map: Maps in Go are unordered. When iterating over a map, the order of elements is not guaranteed:

// for key, value := range scores {
//     fmt.Println(key, value)
// }

//
// ----------------------------------------- Operations ----------------------------------------------
// Widely used to do operations like lookup, update and delete
// In Go, the built-in map type is essentially a hash map.

//
// When we pass the map reference as compared to the variables they will consume low memory
// By-default maps are nil
// map[key_type] value_type {key1: Value1, ...., keyN: ValueN}

func main() {
	var map1 map[string]int
	fmt.Println(map1 == nil) // true

	map2 := map[int]string{1: "1", 2: "2"} // map[key_type] value_type
	map2[3] = "3"

	for k, v := range map2 {
		fmt.Println(k, v)
	}

	map2[4] = "4"
	fmt.Println("Before deleting the value of map :-", map2)

	delete(map2, 4)

	val, ok := map2[4]
	if ok {
		fmt.Println("The value of delete ", val, ok)
	} else {
		fmt.Println("After deleting the value of map :-", map2, ok)
	}
}
