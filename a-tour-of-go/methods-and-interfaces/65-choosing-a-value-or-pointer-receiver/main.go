package main

import (
	"fmt"
	"math"
)

// Vertex is collection fields
type Vertex struct {
	X, Y float64
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.abs())
	v.scale(5)
	fmt.Printf("After scaling: %+v, Abs:%v\n", v, v.abs())
}
