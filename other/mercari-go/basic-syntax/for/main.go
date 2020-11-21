package main

import "fmt"

func main() {
	for i, v := range []int{2, 4, 6, 8, 10} {
		fmt.Println("i, v: ", i, v)
		if i == 3 {
			break
		}
	}
}
