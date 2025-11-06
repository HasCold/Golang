package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Result struct {
	value string
	err   error
}

func worker(jc chan string, wg *sync.WaitGroup, ch chan Result) {
	defer wg.Done()

	for job := range jc {
		time.Sleep(time.Millisecond * 50)

		// Result Channel
		ch <- Result{
			value: job,
			err:   nil,
		}
	}

	fmt.Println("Worker shutting down ...")
}

var wg sync.WaitGroup

var jobs = []string{
	"image_1.png",
	"image_2.png",
	"image_3.png",
	"image_4.png",
	"image_5.png",
	"image_6.png",
	"image_7.png",
	"image_8.png",
	"image_9.png",
	"image_10.png",
	"image_11.png",
	"image_12.png",
	"image_13.png",
	"image_14.png",
	"image_15.png",
	"image_16.png",
	"image_17.png",
	"image_18.png",
	"image_19.png",
	"image_20.png",
}

func main() {
	ch := make(chan Result, 20)
	jobsChannel := make(chan string, len(jobs)) // jobs channel used to send data from main goroutine to workers goroutine.
	totalWorkers := 5

	fmt.Println("Processing Image Started")
	startTime := time.Now()

	// Worker Pool Pattern
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(jobsChannel, &wg, ch)
	}

	// if 20 images has to be processed so 20 goroutines run at the same time but it is not efficient in case of 1000 images which will degrade ur system performance so in this case we have to limit the no. of workers.
	// for _, job := range jobs {
	// wg.Add(1)
	// go worker(job, &wg, ch)
	// }

	// It is a new goroutine which has purpose to wait only
	go func(wg *sync.WaitGroup, ch chan Result) {
		wg.Wait()
		close(ch) // close the channel otherwise cause deadlock
	}(&wg, ch)

	// Send the Jobs to Job Channel
	sendJobs(jobs, jobsChannel)

	fmt.Println("---------------------------------------------------")
	select {
	case val := <-ch:
		fmt.Println("First received from the channel", val)

	default:
		fmt.Println("This is the default value")
	}
	fmt.Println("---------------------------------------------------")

	for res := range ch {
		fmt.Println("----- Job Completed ----- ", res)
		if res.err != nil {
			// Queue - Dead Letter Queue (dlq)
			log.Fatal("error received")
		}
	}

	fmt.Printf("time elapsed %v ms", time.Since(startTime))
}

func sendJobs(jobs []string, jc chan string) {
	for _, j := range jobs {
		jc <- j
	}
	close(jc)
}

// ---------------- Go Concurrency Model ------------------------
// Launch and Wait
// Channels work as a FIFO principle like Queue.
