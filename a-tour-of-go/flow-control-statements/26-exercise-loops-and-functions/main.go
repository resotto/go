package main

import "fmt"

// Sqrt is calculating square root
func Sqrt(x float64) float64 {
	var y float64
	for count, z := 0, 1.0; count < 10; count++ {
		z -= float64((z*z - x) / (2 * z))
		y = z
	}
	return y
}

func main() {
	fmt.Println(Sqrt(144))
}
