// Types

package main

import "fmt"

func main() {
	// Numbers
	fmt.Println("1 + 1 =", 1+1)     // Integers
	fmt.Println("1 + 1 =", 1.0+1.0) // Floating-Point Numbers

	// Strings
	fmt.Println(len("Hello, World")) // Find the length of a string
	fmt.Println("Hello, World"[1])   // Accesses a particular character in the string
	fmt.Println("Hello, " + "World") // Concatenates two strings together

	// Booleans
	fmt.Println(true && true)  // and: true
	fmt.Println(true && false) // and: false
	fmt.Println(true || true)  // or: true
	fmt.Println(true || false) // or: true
	fmt.Println(!true)         // not: false
}
