// Use Case 1: Closing a file
// When opening a file, we should defer file.Close() to ensure the file is closed properly, 
// even if an error occurs or the function exits early.

// Use Case 2: Unlocking a Mutex
// When using sync.Mutex, defer mu.Unlock() ensures that the lock is released, 
// preventing potential deadlocks if the function exits unexpectedly.

// Use Case 3: Closing a Database Connection
// When working with databases, we should defer db.Close() to release the connection 
// once the function completes, avoiding resource leaks.

// Use Case 4: Recovering from a Panic
// If a function panics, using defer with recover() can catch the panic and handle 
// it gracefully instead of crashing the entire program.


package main

import (
	"fmt"
)

func main() {
	// "defer" delays the execution of a function until the surrounding function (main) returns.
	// This line gets executed at the end, after all other statements in main().
	defer fmt.Println("World!") 

	fmt.Println("Hello,")

	fmt.Println("Counting: ")

	// When defer is used inside a loop, it stacks function calls in a LIFO (Last In, First Out) order.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i) // These deferred calls will be executed in reverse order.
	}

	fmt.Println("And Done!")
}

