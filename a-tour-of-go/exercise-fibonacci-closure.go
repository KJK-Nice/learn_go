package main

import "fmt"

// fibonacci is a function that returns
// a function that reutrns an int.
func fibonacci() func() int {
    a, b := 0, 1
    return func() int {
        a, b = b, a + b
        return a
    }
}

func fibonacci2() func() int {
    a := 0
    b := 1
    return func() int {
        temp := a
        a = b
        b = temp + b
        return a
    }
}


func main() {
    f := fibonacci()
    f2 := fibonacci2()
    for i := 0; i < 10; i++ {
        fmt.Printf("f: %f, f2: %f\n", f(), f2())
    }
}
