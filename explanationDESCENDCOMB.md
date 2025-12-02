# Explanation:

package main
- This file is an executable Go program.

import "fmt"
- Imports printing utilities.

func DescendComb() {
    first := true
    - Tracks whether we are printing the first pair, so we avoid an initial comma.

    for a := 99; a >= 10; a-- {
        - Outer loop: a goes from 99 down to 10.
        - a is the first number in the pair.

        for b := a - 1; b >= 0; b-- {
            - Inner loop: b starts one less than a and goes down to 0.
            - Ensures a > b, so pairs descend and no duplicates appear.

            if !first {
                fmt.Print(", ")
                - Print comma before every pair except the first.
            }
            first = false

            fmt.Printf("%02d %02d", a, b)
            - Print the pair using two-digit formatting with leading zeros.
        }
    }

    fmt.Println()
    - End with a newline.
}

func main() {
    DescendComb()
    - Calls the function to print all descending combinations.
}

# Code:

package main

import "fmt"

func DescendComb() {
	first := true

	for a := 99; a >= 10; a-- {
		for b := a - 1; b >= 0; b-- {
			if !first {
				fmt.Print(", ")
			}
			first = false
			fmt.Printf("%02d %02d", a, b)
		}
	}
	fmt.Println()
}

func main() {
	DescendComb()
}
