package main

import (
    "fmt"
    "math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt nagative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
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
	return z, nil
}




func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
