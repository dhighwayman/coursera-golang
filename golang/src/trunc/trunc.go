/*
Write a program which prompts the user to enter a floating point number and
prints the integer which is a truncated version of the floating point number
that was entered. Truncation is the process of removing the digits to the right
of the decimal place.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// prompts asking the number
	var floatingNumber float64
	fmt.Print("Please enter a floating point number (i.e : 123.456) and press ENTER: ")
	fmt.Scan(&floatingNumber)
	fmt.Printf("The truncated number is: %.0f\n", math.Floor(floatingNumber))
}
