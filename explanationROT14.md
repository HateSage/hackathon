Bit-by-bit explanation:

package main
- Declares this file as part of the main executable program.

import (
    "fmt"
)
- Imports the fmt package so we can print characters.

func Rot14(s string) string {
- Defines a function Rot14 that takes a string and returns a rotated string.
- ROT14 means: shift each letter by 14 positions in the alphabet.

    result := []rune{}
    - Creates an empty slice of runes.
    - Using runes allows the function to handle characters safely, especially Unicode.

    for _, r := range s {
        - Loops through every rune (character) in the string.

        if r >= 'A' && r <= 'Z' {
            r = 'A' + (r-'A'+14)%26
            - If r is an uppercase letter:
            - Convert it to a 0–25 range by doing (r - 'A').
            - Add 14 to shift it forward.
            - Use % 26 to wrap around the alphabet.
            - Add 'A' back to turn it into a valid uppercase letter.
        } else if r >= 'a' && r <= 'z' {
            r = 'a' + (r-'a'+14)%26
            - Same logic as above, but for lowercase letters.
        }

        result = append(result, r)
        - Append the transformed (or unchanged) rune to the result slice.
    }

    return string(result)
    - Convert the rune slice back into a string and return it.
}

func main() {
    result := Rot14("Hello! How are You?")
    - Calls Rot14 on the input string.
    - All letters are rotated by 14 places.
    - Non-letters (spaces, punctuation) are unchanged.

    for _, r := range result {
        fmt.Printf("%c", r)
    }
    - Prints each character of the resulting string.

    fmt.Println()
    - Prints a newline after the output.
}
