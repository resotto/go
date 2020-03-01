package main

import (
	"fmt"
	"math"
)

// Abser is interface
type Abser interface {
	abs() float64
}

// MyFloat is float64
type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

// Vertex is collection fields
type Vertex struct {
	X, Y float64
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser
	a = v

	fmt.Println(a.abs())
}
