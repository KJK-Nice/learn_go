package main

import (
    "golang.org/x/tour/tree"
    "fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t.Left != nil {
        Walk(t.Left, ch)
    }
    if t.Right != nil {
        Walk(t.Right, ch)
    }
    ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    fmt.Println("t1: ", t1.String())
    fmt.Println("t2: ", t2.String())
    ch1 := make(chan int, 10)
    ch2 := make(chan int, 10)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for i := 0; i < 10; i++ {
        t1v := <-ch1
        t2v := <-ch2
        fmt.Println(t1v, t2v)
        if t1v != t2v {
            return false
        }
    }
    return true
}

func main () {
    // test := tree.New(1)
    // fmt.Println(Same(test, test))
    fmt.Println(Same(tree.New(1), tree.New(1)))
}

