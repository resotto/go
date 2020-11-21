package main

import "fmt"

// Vertex is collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
