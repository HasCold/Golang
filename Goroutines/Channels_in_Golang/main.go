package main

import (
	"fmt"
)

// Channnels
// In Go, channels are used for communication between goroutines. They allow goroutines to synchronize and exchange data safely.
// When you are working with multiple goroutines then you need to work with channels
// Suppose one goroutine is depend on other goroutine so you have to fullfil the need of that goroutine via channels

// Syntax
// var Channel_name chan Type
// Channels using make short-hand method
// channel_name := make(chan Type)

func main() {
	//	---------------------------- Declaring Channel --------------------------------------
	var firstChan chan int
	// firstChan = make(chan int, 10)  // Initialize the buffered channel
	// firstChan := make(chan int) // // Initialize the Unbuffered channel

	fmt.Println("Value of the channel :-", firstChan)    // 0xc00001c0c0
	fmt.Printf("Type of the channel '%T' \n", firstChan) // It’s particularly useful for debugging or when you need to understand the type of a variable at runtime. //  Type :- chan int

	//
	//	---------------------------- Assigning Value to Channel --------------------------------------
	// chan1 <- value  // Value transfer or send to the channel 1
	// a := <- chan1  // Channel value store into the "a" variable
	// <- chan1			// Print the channel value

	//
	// ----------------------------- Working with Channel ------------------------------
	ch := make(chan int)
	fmt.Println("Hello from main")

	// Basically controller doesn't wait for the goroutine
	go multiplyWithChannel(ch) // This line starts a new goroutine that executes the multiplyWithChannel function. The ch channel is passed as an argument to this function.
	ch <- 9                    // This line sends the integer value 9 to the ch channel. This operation is blocking, meaning the main goroutine will wait until the value is received by another goroutine before it proceeds. Since the main goroutine is sending a value to the channel, it will pause until the multiplyWithChannel function receives that value.
	fmt.Println("Byte from main")

	//
	// ----------------------------- Check and Close with Ok Syntax ------------------------------
	fmt.Println("------------------------------------------------------------------------------------")
	ch1 := make(chan int)
	close(ch1) // Close is the built-in method to close the channels
	elem, ok := <-ch1
	fmt.Println("Hello from Close Syntax", elem, ok)

	//
	// ----------------------------- Loops On Channel ------------------------------
	fmt.Println("------------------------------------------------------------------------------------")
	c := make(chan string) // Basically channels are used to communicate between goroutines
	go initStrings(c)
	for {
		resp, ok := <-c // Get the value from the channel c
		if ok == false {
			fmt.Println("Channel Close", ok)
			break
		}
		fmt.Println("Channel Open", resp, ok)
	}

	//
	// ----------------------------- Length Of Channel ------------------------------
	fmt.Println("------------------------------------------------------------------------------------")
	myChnl := make(chan string, 4) // Buffered Channel has their length while Unbuffered Channel doesn't have their length
	myChnl <- "abc"
	myChnl <- "Hasan"
	myChnl <- "M."
	myChnl <- "Hello"
	// myChnl <- "World"
	fmt.Println("Length of the channel :- ", len(myChnl))

	//
	// ----------------------------- Channel with Range Loop ------------------------------
	fmt.Println("------------------------------------------------------------------------------------")
	testCh := make(chan string)
	go func() {
		testCh <- "Value1"
		testCh <- "Value2"
		close(testCh)
	}()

	for res := range testCh { // range loop jab tak channel ke andr kuch ha tu wo uski value ko access krta rahega
		fmt.Println(res)
	}
}

func multiplyWithChannel(val chan int) {
	fmt.Println(100 * <-val) // Multitply the value of channel
}

func initStrings(chnl chan string) {
	for v := 0; v < 3; v++ {
		chnl <- "From Hasan"
	}
	close(chnl)
}
