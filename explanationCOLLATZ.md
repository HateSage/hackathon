# Explanation:

package main
- Makes this file an executable Go program.

import "fmt"
- Lets us print text using fmt.Println.

func CollatzCountdown(start int) int {
    if start <= 0 {
        return -1
    }
    - Collatz requires a positive integer. Return -1 for invalid input.

    steps := 0
    - Counter for how many operations we perform.

    for start != 1 {
        if start%2 == 0 {
            start = start / 2
            - Even number: divide by 2.
        } else {
            start = 3*start + 1
            - Odd number: apply 3n + 1.
        }
        steps++
        - Count the step.
    }

    return steps
    - Returns the total number of steps it took to reach 1.
}

func main() {
    steps := CollatzCountdown(12)
    - Call the function with 12.

    fmt.Println(steps) // Expected output: 9
    - Print how many steps it took (9).
}

# Code:

package main

import "fmt"

func CollatzCountdown(start int) int {
    if start <= 0 {
        return -1
    }

    steps := 0
    for start != 1 {
        if start%2 == 0 {
            start = start / 2
        } else {
            start = 3*start + 1
        }
        steps++
    }

    return steps
}

func main() {
    steps := CollatzCountdown(12)
    fmt.Println(steps) // Expected output: 9
}
