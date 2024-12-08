package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 0, 1
    for {
        select {
        case c <- x:
            x, y = y, x+y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}

// please explain how this code works
//
// The fibonacci function is a generator that generates Fibonacci numbers on the channel c.
// The main function starts a goroutine that consumes the first 10 values from the channel c and then signals on the channel quit to terminate the Fibonacci generator.
// The select statement in the Fibonacci generator allows it to exit when the quit channel is closed.
//
// The output of this program is:
//  0
//  1
//  1
//  2
//  3
//  5
//  8
//  13
//  21
//  34
//  quit
//
// Note that the quit message is printed after the first 10 Fibonacci numbers.
//
// The select statement in the fibonacci function works like a switch statement, but the case statements are channel operations.
// The select statement blocks until one of its cases can proceed, then it executes that case.
// If multiple cases can proceed, select chooses one at random.
//
// The select statement in the fibonacci function has two cases:
//  the first case sends the Fibonacci number on the channel c,
//  and the second case receives a value from the channel quit.
// The select statement blocks until one of the cases can proceed.
// If the quit channel is closed, the second case will proceed, and the function will return.
//
// The main function starts a goroutine that consumes the first 10 Fibonacci numbers from the channel c.
// After consuming the first 10 numbers, the goroutine sends a value on the quit channel to signal the Fibonacci generator to exit.
// The main function then waits for the Fibonacci generator to exit by calling fibonacci(c, quit).


