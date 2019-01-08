package main

import "fmt"

type Vertex struct {
    X int
    Y int
}

var (
    v1 = Vertex{1,2}
    v2 = Vertex{X: 3}
    v3 = Vertex{}
    p = &Vertex{X: 10, Y: 12}
)
func main() {
    fmt.Println(v1, v2, v3, p)
}

