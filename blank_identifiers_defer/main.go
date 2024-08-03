package main

import "fmt"

func main() {
	// Blank Identifier
	a, _, n := identi() // Output :-  1 <nil>
	fmt.Println(a, n)

	// Range Loop
	arr := []int{0, 2, 3, 5}
	for _, val := range arr {
		fmt.Printf("The value is %v \n", val)
	}

	// Defer Keyword :- defer keyword mainly used for file closing or connection closing just to aviod the miscellaneous activity e.g. For database connection closing
	// When used defer keyword with any function they must delay for sometime
	// defer follows the LIFO order (Last In First Out)
	defer func() {
		fmt.Println("Invoked Defer Closing")
	}()

	fmt.Println("I am done")

}

func identi() (int, int, error) {
	return 1, 2, nil
}
