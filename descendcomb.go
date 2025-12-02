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
