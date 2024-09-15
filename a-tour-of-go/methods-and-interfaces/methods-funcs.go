package main

import (
    "fmt"
    "math"
)

type Vertex struct {
    X, Y float64
}

// a regular function
func Abs(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// declare a method on non-struct types
type MyFloat float64

// cannot declare a method with a receiver
// whose type is defined in another package
func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main() {
    v := Vertex{3, 4}
    fmt.Println(Abs(v))

    f := MyFloat(-math.Sqrt2)
    fmt.Println(f.Abs())
}
