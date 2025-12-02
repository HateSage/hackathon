package main

import "fmt"

func Unmatch(a []int) int {
	counts := make(map[int]int)

	for _, n := range a {
		counts[n]++
	}

	for n, c := range counts {
		if c%2 != 0 {
			return n
		}
	}

	return -1
}

func main() {
	a := []int{1, 2, 3, 1, 2, 3, 4, 4, 4}
	unmatch := Unmatch(a)
	fmt.Println(unmatch) // Expected: 4
}
