package main

import (
	"fmt"
	"math"
)

// Vertex is collection fields
type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func absFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())
	fmt.Println(absFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.abs())
	fmt.Println(absFunc(*p))
}
