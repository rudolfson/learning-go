package main

import "fmt"

func main() {
    primes := [6]int{2, 3, 5, 7, 11, 13}

    var slice []int = primes[1:4]
    fmt.Println(slice)
    fmt.Println(len(primes))
}
