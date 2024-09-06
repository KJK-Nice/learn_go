package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

// struct literals
var (
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0
    p1  = &Vertex{1, 2} // has type *Vertex
)

func main() {
    v := Vertex{1, 2}
    p := &v // also written as var p *Vertex = &v
    p.X = 1e9 // also written as (*p).X = 1e9
    fmt.Println(v, p)

    // struct literals
    fmt.Println(v1, p1, v2, v3)
}
