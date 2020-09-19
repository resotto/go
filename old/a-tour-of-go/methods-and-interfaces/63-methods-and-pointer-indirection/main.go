package main

import "fmt"

// Vertex is collection fields
type Vertex struct {
	X, Y float64
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.scale(2)
	scaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.scale(3)
	scaleFunc(p, 8)

	fmt.Println(v, p)
}
