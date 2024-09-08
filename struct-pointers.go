package main

import "fmt"

type Vertex struct {
    X int
    Y int
}


func main() {
    v := Vertex{1, 2}
    p := &v // also written as var p *Vertex = &v
    p.X = 1e9 // also written as (*p).X = 1e9
    fmt.Println(v, p)
}
