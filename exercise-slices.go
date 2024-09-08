package main

import "golang.org/x/tour/pic"


func Pic(dx, dy int) [][]uint8 {
    x := make([][]uint8, dx)
    for i := range x {
        x[i] = make([]uint8, dy)
        for j := range x[i] {
            x[i][j] = uint8(j*i)
        }
    }

    return x
}

func main() {
    pic.Show(Pic)
}
