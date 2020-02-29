package main

import (
	"fmt"
	"math"
)

func main() {
	//	fmt.Println(math.pi) // Error
	fmt.Println(math.Pi) // Capital letter is an exported name and is only able to be imported
}
