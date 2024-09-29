package main

import "fmt"

type I interface {
    M()
}

// Calling a method on a nil interface is a run-time error
// because there is no type inside the inteface tuple
// to indicate which concrete method to call.

func main() {
    var i I
    describe(i)
    i.M()
}

func describe(i I) {
    fmt.Printf("(%v, %T)\n", i, i)
}
