package main

import "fmt"

func LoafOfBread(str string) string {
	letters := ""
	for _, ch := range str {
		if ch != ' ' {
			letters += string(ch)
		}
	}

	if len(letters) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	count := 0

	for _, ch := range letters {
		result += string(ch)
		count++
		if count == 5 {
			result += " "
			count = 0
		}
	}

	return result
}

func main() {
	fmt.Print(LoafOfBread("deliciousbread"))
	fmt.Print(LoafOfBread("This is a loaf of bread"))
	fmt.Print(LoafOfBread("loaf"))
}
