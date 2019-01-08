package main

import (
    "fmt"
    "math"
)

const ITERATIONS = 10
const DIFF = 0.00000001

func Sqrt(x float64) float64 {
    z := 1.0
    prev := z - 10*DIFF
    for i := 0; i < ITERATIONS && math.Abs(z - prev) > DIFF; i++ {
        prev = z
        z -= (z*z - x) / (2*z)
        fmt.Println("z is", z)
    }
    return z
}

func main() {
    var x = 2.0
    fmt.Printf("sqrt(%g)\n", x)
    fmt.Println(Sqrt(x))
    fmt.Println(math.Sqrt(x))
    x = 3.0
    fmt.Printf("sqrt(%g)\n", x)
    fmt.Println(Sqrt(x))
    fmt.Println(math.Sqrt(x))
    x = 300.0
    fmt.Printf("sqrt(%g)\n", x)
    fmt.Println(Sqrt(x))
    fmt.Println(math.Sqrt(x))
    x = 303240.0
    fmt.Printf("sqrt(%g)\n", x)
    fmt.Println(Sqrt(x))
    fmt.Println(math.Sqrt(x))
}

