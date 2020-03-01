package main

import (
	"fmt"
	"math"
)

// Vertex is collection fields
type Vertex struct {
	X, Y float64
}

// Abs calculates v.X*v.X + v.Y*v.Y of Vertex
func (v Vertex) Abs() float64 {
	fmt.Println("v when Abs: ", v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale multiply each value of Vertex with specified value
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("v when Scale: ", v)
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
