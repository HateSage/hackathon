Bit-by-bit explanation:

package main
- Declares this file as part of the main Go program.

import "fmt"
- Imports the fmt package for printing to the console.

func Unmatch(a []int) int {
- Defines a function Unmatch.
- It takes a slice of integers and returns an integer.
- The goal: find the number that appears an odd number of times.

    counts := make(map[int]int)
    - Creates a map where:
      key   = the number in the slice
      value = how many times it has appeared

    for _, n := range a {
        counts[n]++
    }
    - Loops through each number n in the slice.
    - For each number, increase its count in the map by 1.

    for n, c := range counts {
        if c%2 != 0 {
            return n
        }
    }
    - Loops through the map.
    - If a number’s count is odd (c % 2 != 0), return that number.
    - This is the “unmatched” number.

    return -1
    - If no number appears an odd number of times, return -1.
    - This means no unmatched value exists.
}

func main() {
    a := []int{1, 2, 3, 1, 2, 3, 4, 4, 4}
    - Creates a slice of integers with some repeated values.

    unmatch := Unmatch(a)
    - Calls Unmatch to find the number with an odd count.

    fmt.Println(unmatch) // Expected: 4
    - Prints the result (4), because:
      1 appears 2 times
      2 appears 2 times
      3 appears 2 times
      4 appears 3 times (odd)
}
