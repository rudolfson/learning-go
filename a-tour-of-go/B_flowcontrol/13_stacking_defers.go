package main

import "fmt"

func main() {
    fmt.Println("counting")

    for i := 0; i < 3; i++ {
        defer fmt.Println("Printing", i)
    }

    fmt.Println("done")
}
