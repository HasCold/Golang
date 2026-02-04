package main

func main() {

}

// Goroutines in Go are considered **lightweight threads** when compared to traditional threads used in other programming languages like Java, C++, or Python. There are several key reasons for this, primarily related to their design, scheduling, memory usage, and management.
// Goroutines will be managed by go runtime scheduler.

// ### Key Reasons Why Goroutines are Lightweight

// #### 1. **Smaller Stack Size**:
//    - **Goroutines** start with a very small stack size, typically around **2KB**. This stack can grow and shrink dynamically as needed. If a goroutine requires more stack space, the Go runtime automatically allocates more memory.
//    - **Traditional threads** (OS threads) generally start with a much larger stack size (e.g., **1MB**), even if they don’t use most of it. This results in more memory consumption per thread.

//    - **Advantage of Goroutines**: Since they start with a smaller stack, you can create a large number of goroutines without exhausting system memory, whereas creating too many OS threads can lead to resource exhaustion.

// #### 2. **User-Space Scheduling (M:N Model)**:
//    - **Goroutines** are scheduled by the Go runtime, not by the operating system. This is called **user-space scheduling** or **cooperative multitasking**. Go uses an **M:N scheduling model**, meaning **M goroutines** are multiplexed onto **N OS threads**. The Go scheduler switches between goroutines without needing OS intervention.
//    - **Traditional threads** are scheduled by the operating system’s kernel, which involves more overhead because the OS must manage context switching, which can be slow and resource-heavy.

//    - **Advantage of Goroutines**: The Go runtime can switch between goroutines more efficiently than the OS can switch between threads. This reduces context-switching overhead and allows for thousands or even millions of goroutines to run concurrently with minimal overhead.

// #### 3. **No Context Switching in the Kernel**:
//    - In **traditional threads**, when switching between threads, the operating system has to perform **context switching** (saving and restoring the state of CPU registers, memory, etc.) at the kernel level, which is expensive.
//    - **Goroutines** avoid this because context switching is handled by the Go runtime in user space, making it more efficient.

//    - **Advantage of Goroutines**: They avoid the costly kernel-level context switches, leading to faster and more efficient multitasking.

// #### 4. **Cheap to Create and Destroy**:
//    - **Goroutines** are extremely cheap to create and destroy because of their smaller initial stack size and the lightweight nature of the Go scheduler.
//    - **Traditional threads** are more expensive to create and destroy due to the involvement of the OS in thread management (allocating memory, setting up thread context, etc.).

//    - **Advantage of Goroutines**: Creating thousands or even millions of goroutines is feasible with little memory and processing overhead, while creating an equivalent number of OS threads would be resource-intensive and impractical.

// #### 5. **Concurrency Model with Channels**:
//    - **Goroutines** are designed with Go's concurrency model in mind, which uses **channels** for communication and synchronization. Channels provide a simple and effective way to pass data between goroutines without having to deal with lower-level synchronization primitives like mutexes and condition variables.
//    - **Traditional threads** require the use of locks (e.g., mutexes) or other synchronization primitives to prevent race conditions, which introduces complexity and potential performance bottlenecks.

//    - **Advantage of Goroutines**: Go’s concurrency model is easier to use and more efficient, especially for managing communication between goroutines, avoiding common problems like deadlocks.

// #### 6. **Garbage Collection**:
//    - **Goroutines** benefit from Go's **garbage collector**, which helps manage memory automatically. The runtime automatically reclaims memory that is no longer in use, reducing the developer's need to manually manage memory.
//    - In many other languages, you need to be more careful with memory management when using threads, especially with respect to memory leaks or dangling pointers.

//    - **Advantage of Goroutines**: Memory management is simpler and less error-prone in Go compared to languages where manual memory management is required.

// ### Comparison Summary

// | Feature                    | Goroutines (Go)                                        | Traditional Threads (OS-managed)                     |
// |----------------------------|-------------------------------------------------------|-----------------------------------------------------|
// | **Stack Size**              | Starts at ~2KB, grows dynamically                      | Starts at ~1MB, fixed by default                     |
// | **Scheduling**              | User-space scheduling (Go runtime)                     | OS-based scheduling                                 |
// | **Context Switching**       | Lightweight, managed by Go runtime (user space)        | Expensive, involves OS kernel                       |
// | **Creation Cost**           | Very low                                               | High (OS overhead)                                  |
// | **Memory Usage**            | Low (small initial stack, dynamic growth)              | High (large initial stack)                          |
// | **Concurrency Model**       | Channels for communication                            | Locks, mutexes, semaphores for synchronization       |
// | **Number of Concurrent Tasks** | Easily supports millions of goroutines               | Limited by OS resources, typically tens of thousands |
// | **Garbage Collection**      | Automatic (Go runtime)                                | Depends on the language and environment              |

// ### Conclusion:
// - **Goroutines** are much more lightweight compared to traditional threads. Their small initial memory footprint, dynamic stack growth, user-space scheduling, and minimal context-switching overhead make them ideal for large-scale concurrent programs.
// - This allows Go programs to efficiently handle **thousands or millions of concurrent tasks**, something that would be difficult and resource-intensive with traditional OS threads.
