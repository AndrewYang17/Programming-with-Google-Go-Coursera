// Write a program which prompts the user to enter a floating point number and
// prints the integer which is a truncated version of the floating point number
// that was entered. Truncation is the process of removing the digits to the
// right of the decimal place.

package main

import "fmt"

func main() {
	var getValue float64
	fmt.Println("Please enter a float value: ")
	_, _ = fmt.Scan(&getValue)
	fmt.Println("The value after truncating is:", int(getValue))
}
