package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x < 0 {
		return math.NaN() // Return NaN for negative inputs
	}

	z := 1.0
	epsilon := 1e-12 // Define the precision of the result
	count := 0
	for {
		nextZ := z - (z*z-x)/(2*z) // Update the value of z
		if math.Abs(nextZ-z) < epsilon {
			break // Stop when the difference is small enough
		}
		z = nextZ
		count++
	}
	fmt.Println("count", count)
	return z
}

func main() {
	fmt.Println(Sqrt(2))      // Your custom square root function
	fmt.Println(math.Sqrt(2)) // Built-in square root function
}

