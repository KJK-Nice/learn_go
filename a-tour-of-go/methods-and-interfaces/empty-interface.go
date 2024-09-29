package main

import "fmt"

func main() {
    // the interface type that specifies zero methods
    // is knwon as the empty inteface
    var i interface{}
    describe(i)

    // an empty interface may hold values of any type.
    i = 42
    describe(i)

    // Empty interface are used by code that handles
    // values of unknown type.
    i = "hello"
    describe(i)
}

func describe(i interface{}) {
    fmt.Printf("(%v, %T)\n", i, i)
}
