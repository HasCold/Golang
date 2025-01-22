package main

func main() {

}

// In Go, the decision about whether a variable is allocated on the **heap** or the **stack** depends on the **scope of the variable** and how it is used. Let's break down your example:

// ### **Declaration**
// ```go
// var a = 5
// ```

// 1. **What happens here?**
//    - `a` is a global variable if declared outside any function or block.
//    - If declared inside a function, it becomes a local variable.

// ### **Heap vs. Stack Allocation**
// #### **Stack Allocation**
// - **Stack** is used for variables with limited lifetime, such as those declared inside a function or block.
// - Variables stored on the stack are automatically cleaned up when the function exits (stack unwinding).
// - If `a` is declared inside a function and its address is not returned or used outside the function, it will typically be allocated on the stack.

// Example:
// ```go
// func main() {
//     var a = 5 // Likely allocated on the stack
//     fmt.Println(a)
// }
// ```

// #### **Heap Allocation**
// - **Heap** is used for variables that need to live beyond the function where they were created.
// - If you take the address of a variable (using `&a`) and that address is used outside the function, the variable is moved to the heap so it persists.

// Example:
// ```go
// func getPointer() *int {
//     var a = 5
//     return &a // `a` escapes to the heap because its pointer is returned
// }

// func main() {
//     p := getPointer()
//     fmt.Println(*p)
// }
// ```

// In this case, Go's compiler performs **escape analysis** and determines that `a` must be moved to the heap because it needs to exist even after `getPointer` exits.

// ### **Global Variables**
// - If `var a = 5` is declared outside any function, it is a **global variable** and is always allocated on the heap.
//   ```go
//   var a = 5 // Allocated on the heap

//   func main() {
//       fmt.Println(a)
//   }
//

// ### **Conclusion**
// - **Stack**: Used for local variables with limited lifetime and no references outside the function.
// - **Heap**: Used for variables that escape their defining scope (e.g., accessed via pointers) or are global variables.

// Go's compiler performs **escape analysis** to decide where a variable should be allocated, optimizing for performance and memory usage. You don't explicitly control this, but understanding these rules helps write more efficient code.
