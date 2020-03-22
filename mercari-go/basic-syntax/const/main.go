package main

import "fmt"

func main() {
	// infinite accuracy because const doen't have type
	const n = 10000000000000000000 / 10000000000000000000
	var m = n  // type of m is int
	println(m) // 1

	const helloWorld = "Hello, World"
	println(helloWorld)

	const (
		a = 1 + 2
		b
		c
	)
	fmt.Println(a, b, c)
}
