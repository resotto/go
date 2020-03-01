package main

import "fmt"

// I is interface
type I interface {
	m()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I
	describe(i)
	i.m() // error!
}
