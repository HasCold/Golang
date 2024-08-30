package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("------------------- Maths in Golang -------------------")

	var myNumberOne int = 12
	var myNumberTwo float64 = 4.5

	fmt.Println("Ths sum is :", myNumberOne+int(myNumberTwo))

	// Math Random Number

	// Random Number ; rand.Intn -->> Random Integer
	// rand.Seed(time.Now().UnixNano())
	// // random number will exclusively mention to the function
	// i := rand.Intn(5) + 1 // Generates a random number between 0 and 4, then add 1
	// fmt.Println(i)

	// Cryptography ; Crypto Random Number
	myRandInt, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandInt)
}
