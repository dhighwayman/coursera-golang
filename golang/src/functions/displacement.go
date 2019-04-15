package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//GenDisplaceFn return a function which computes displacement as a function of time, assuming the given values acceleration, initial velocity, and initial displacement.
func GenDisplaceFn(a, v0, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return a/2*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	a, v0, s0 := 0.0, 0.0, 0.0
	fmt.Println("Please enter acceleration (float)")
	fmt.Scan(&a)
	fmt.Println("Please enter speed (float)")
	fmt.Scan(&v0)
	fmt.Println("Please enter initial displacement (float)")
	fmt.Scan(&s0)
	fn := GenDisplaceFn(a, v0, s0)
	for {
		fmt.Println("Enter a time value as a floating point number or x to quit")
		var input string
		fmt.Scan(&input)
		if strings.ToLower(input) == "x" {
			fmt.Print("Quitting...\n")
			return
		} else if time, err := strconv.ParseFloat(input, 64); err == nil {
			fmt.Printf(">%f\n", fn(time))
		} else {
			fmt.Printf("%s is not a valid float", input)
		}

	}

}
