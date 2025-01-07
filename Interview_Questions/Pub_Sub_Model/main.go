package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Message struct {
	Content string
	ID      int
}

var wg = &sync.WaitGroup{}

func main() {
	messageChannel := make(chan Message, 10)

	fmt.Println("Register New Consumer ID Separated by Space :- ")
	// Create a scanner to read input from the command line
	scanner := bufio.NewScanner(os.Stdin)

	// Read the input
	ok := scanner.Scan()
	if !ok {
		fmt.Println("Error reading from stdin")
		return
	}
	input := scanner.Text()

	// Split the input string into a slice of string
	strSlice := strings.Fields(input)

	// Convert the slice string into the slice integer
	intSlice := make([]int, len(strSlice))
	for _, str := range strSlice {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Panic(err)
		}

		intSlice = append(intSlice, num)
	}

	wg.Add(2)

	go producer(messageChannel, wg)
	go consumer(messageChannel, wg, intSlice)

	select { // select handles the first incoming channels
	case id := <-messageChannel:
		fmt.Printf("Received Signal From Channel %v \n", id)
	default:
		fmt.Println("No Signal Received")
	}

	wg.Wait()
	fmt.Println("All processing finished !")
}

func producer(ch chan Message, wg *sync.WaitGroup) {
	fmt.Println("Producer Working Starts...")
	time.Sleep(time.Second * 2)

	counter := 1

	for counter < 10 {
		message := Message{
			Content: fmt.Sprintf("Message: Hello, Ayesha %v", counter),
			ID:      counter - 1,
		}
		ch <- message
		counter++
	}

	close(ch)
	defer wg.Done()
}

func consumer(ch chan Message, wg *sync.WaitGroup, userId []int) {
	userIdMap := make(map[int]bool)

	for _, num := range userId {
		userIdMap[num] = true
	}

	for msg := range ch {
		if userIdMap[msg.ID] {
			fmt.Printf("Consumed Message :- %v and corresponding ID :- %v \n", msg.Content, msg.ID)
			delete(userIdMap, msg.ID) // delete the id after processing to avoid the duplication
		}
	}

	defer wg.Done()
}
