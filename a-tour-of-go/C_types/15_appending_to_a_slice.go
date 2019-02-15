package main

import "fmt"

func main() {
    var s []int
    printSlice(s)

    // append works on nil slices
    s = append(s, 0)
    printSlice(s)

    // the slice grows as needed
    s = append(s, 1)
    printSlice(s)

    // add more than one element
    s = append(s, 2, 3, 4)
    printSlice(s)

    s = append(s, 5, 6)
    printSlice(s)
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
