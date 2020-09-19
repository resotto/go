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
func Abs(v Vertex) float64 {
	fmt.Println("v when Abs: ", v)
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale multiplies each value of Vertex
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
	fmt.Println("v when Scale: ", v)
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
