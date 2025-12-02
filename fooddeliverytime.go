package main

import "fmt"

func FoodDeliveryTime(order string) int {
	switch order {
	case "burger":
		return 15
	case "chips":
		return 10
	case "nuggets":
		return 12
	default:
		return 404
	}
}

func main() {
	fmt.Println(FoodDeliveryTime("burger"))  // 15
	fmt.Println(FoodDeliveryTime("chips"))   // 10
	fmt.Println(FoodDeliveryTime("nuggets")) // 12

	total := FoodDeliveryTime("burger") +
		FoodDeliveryTime("chips") +
		FoodDeliveryTime("nuggets")

	fmt.Println(total) // 37
}
