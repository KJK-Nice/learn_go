package main

import "fmt"

func main() {
    a := make([]int, 5)
    printSlice("a", a) // [0 0 0 0 0] len=5 cap=5

    b := make([]int, 0, 5) // this line specify cap by passing 3rd argrument
    printSlice("b", b) // [] len=0 cap=5

    c := b[:2]
    printSlice("c", c) // [0 0] len=2 cap=5

    d := c[2:5]
    printSlice("d", d) // [0 0 0] len=3 cap=3
}

func printSlice(s string, x []int) {
    fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
