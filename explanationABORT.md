# Explanation:

package main
- Defines the main package, meaning this file is part of a runnable Go program.

import "fmt"
- Imports the fmt package for printing output.

func Abort(a, b, c, d, e int) int {
    return (a + b + c + d + e) / 5
}
- Defines a function named Abort.
- It takes five integers as input.
- It returns the average of the five integers using integer division.

func main() {
    middle := Abort(2, 3, 8, 5, 7)
    fmt.Println(middle) // Expected output: 5
}
- main() is the entry point of the Go program.
- Calls Abort with five numbers.
- Stores the result in 'middle'.
- Prints the result (5).

# Code:

package main

import "fmt"

func Abort(a, b, c, d, e int) int {
    return (a + b + c + d + e) / 5
}

func main() {
    middle := Abort(2, 3, 8, 5, 7)
    fmt.Println(middle) // Expected output: 5
}
