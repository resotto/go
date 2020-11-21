package main

import (
	"fmt"
	"math"
)

// Vertex is collection fields
type Vertex struct {
	X, Y float64
}

// Abs : Calculate Sqrt of v.X*v.X + v.Y*v.Y of Vertex
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
