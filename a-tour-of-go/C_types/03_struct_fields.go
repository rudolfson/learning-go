package main

import "fmt"

type Vertex struct {
    X int
    Y int
    z int
}
func main() {
    v := Vertex{1, 2, 3}
    v.X = 5
    fmt.Println(v)
    fmt.Println(v.z)
}

