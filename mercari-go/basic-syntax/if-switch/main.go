package main

import "fmt"

func main() {
	if a := 0; a > 0 {
		fmt.Println(a)
	} else {
		fmt.Println(a + 1)
	}

	b := 23
	switch b {
	case 23:
		fmt.Println("b is 23")
		fallthrough
	case 1, 2:
		fmt.Println("b is 1 or 2")
	default:
		fmt.Println("default")
	}

	switch {
	case b == 23:
		fmt.Println("b is 23 again")
	}
}
