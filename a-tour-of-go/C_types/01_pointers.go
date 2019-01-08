package main

import "fmt"

func main() {
    var p *int
    i := 42
    p = &i

    fmt.Println("i is", i)
    fmt.Println("p is", *p)
    *p = 21
    fmt.Println("i is", i)
    fmt.Println("p is", *p)
}
