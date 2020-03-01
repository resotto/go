package main

import (
	"fmt"
	"math"
)

// Vertex is colleciton fields
type Vertex struct {
	X, Y float64
}

// Abs calculates v.X*v.X + v.Y*v.Y of Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
