// -------------------------------- Deadlock Explanation ------------------------------------
// In Go, a deadlock occurs when two or more goroutines are waiting on each other to finish some action, but none of them can proceed because they are all waiting. In the case of channels, a deadlock can happen if you try to send or receive values when there is no corresponding receiver or sender available.

package main

import "fmt"

func main() {
	ch := make(chan int) // This line creates a new channel ch.
	fmt.Println("Hello from main")

	// Attempt to send a value to the channel before the goroutine starts
	ch <- 9 // Here, the main goroutine tries to send the value 9 to the channel. However, there is no receiver at this point because the goroutine has not yet started. Since the channel is unbuffered, the send operation will block until a receiver is available.

	// Start the goroutine
	go multiplyWithChannel(ch) // This line starts a new goroutine that waits to receive from the channel ch. But, by the time this goroutine starts, the main goroutine is already blocked waiting to send the value 9. Thus, neither the main goroutine nor the new goroutine can proceed, resulting in a deadlock.

	fmt.Println("Byte from main")
}

func multiplyWithChannel(val chan int) {
	fmt.Println(100 * <-val)
}

// ----------------------------------------- Unbuffered vs. Buffered Channels ---------------------------------
// Unbuffered Channels: An unbuffered channel requires both a sender and a receiver to be ready for the operation to complete. If you attempt to send to an unbuffered channel without a receiver ready, or receive from an unbuffered channel without a sender, it results in a deadlock.

// Buffered Channels: Buffered channels can hold a certain number of values before blocking. If you use a buffered channel with sufficient capacity, the send operation will only block if the buffer is full. However, in the provided example, the channel is unbuffered.

// To avoid deadlock with a buffered channel, you can use the following approach:

func avoidBlocking() {
	ch := make(chan int, 1) // Create a buffered channel with capacity 1
	fmt.Println("Hello from main")

	ch <- 9 // Send value to the channel

	go multiplyWithChannel1(ch) // Start the goroutine

	fmt.Println("Byte from main")
}

func multiplyWithChannel1(val chan int) {
	fmt.Println(100 * <-val)
}

// In this example, the channel has a buffer size of 1, so the send operation does not block. The value 9 can be sent to the channel, and then the goroutine can be started and receive the value.
