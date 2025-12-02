package main

import (
	"fmt"
)

func Rot14(s string) string {
	result := []rune{}

	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = 'A' + (r-'A'+14)%26
		} else if r >= 'a' && r <= 'z' {
			r = 'a' + (r-'a'+14)%26
		}
		result = append(result, r)
	}

	return string(result)
}

func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		fmt.Printf("%c", r)
	}
	fmt.Println()
}
