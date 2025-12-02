Bit-by-bit explanation:

package main
- This declares the main package. Any Go program that produces an executable must have a main package.

import "fmt"
- Imports the fmt package so the program can print text.

func FoodDeliveryTime(order string) int {
- Defines a function named FoodDeliveryTime.
- It accepts one parameter: order (a string).
- It returns an integer that represents the delivery time in minutes.

    switch order {
    - A switch statement is used to match the value of 'order' against different cases.

    case "burger":
        return 15
        - If order is "burger", the function returns 15 minutes.

    case "chips":
        return 10
        - If order is "chips", return 10 minutes.

    case "nuggets":
        return 12
        - If order is "nuggets", return 12 minutes.

    default:
        return 404
        - If the order doesn't match any known food, return 404.
        - 404 is used here as a “not found” value.

}

func main() {
- The main entry point of the program.

    fmt.Println(FoodDeliveryTime("burger"))
    - Calls FoodDeliveryTime with "burger" and prints 15.

    fmt.Println(FoodDeliveryTime("chips"))
    - Prints 10.

    fmt.Println(FoodDeliveryTime("nuggets"))
    - Prints 12.

    total := FoodDeliveryTime("burger") +
        FoodDeliveryTime("chips") +
        FoodDeliveryTime("nuggets")
    - Calculates the total of all three delivery times.
    - 15 + 10 + 12 = 37.

    fmt.Println(total)
    - Prints 37.
}
