package main

// The difference between Concurrency and Parallelism

// Concurrency :-
// . Suppose a Task have multiple operations so in concurrency will run our operations and also manages the operations
// . One task is completed by multiple processes individually / One task is depend on multiple processes

// Characteristics:
// . Task Management: Concurrency involves managing multiple tasks by switching between them, often so quickly that it appears they are running simultaneously.
// . Single-Core: Concurrency can occur on a single-core processor where multiple tasks are interleaved.
// . Goroutines: In Go, concurrency is achieved using goroutines, which are lightweight threads managed by the Go runtime.

// ---------------------------------------------------------------------------------------------------------------------

// Parallelism :-
// . Suppose a Task have multiple operations so in parallelism will run our operations but not manages the operations
// . One task is completed by multiple processes simultaneously and that will also not manages the task.

// Characteristics:
// . Simultaneous Execution: Parallelism requires hardware with multiple cores or processors to truly run tasks at the same time.
// . Multi-Core: Parallelism can only be achieved on multi-core processors.
// . Performance: Parallelism can significantly increase the performance of a program by dividing the workload across multiple processors.

func main() {

}
