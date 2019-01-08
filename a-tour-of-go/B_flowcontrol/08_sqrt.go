package main

import (
    "fmt"
    "math"
)

const ITERATIONS = 10

func Sqrt(x float64) float64 {
    var z float64 = 1.0
    for i := 0; i < ITERATIONS; i++ {
        z -= (z*z - x) / (2*z)
        fmt.Println("z is", z)
    }
    return z
}

func main() {
    var x = 3.0
    fmt.Println(Sqrt(x))
    fmt.Println(math.Sqrt(x))
}

