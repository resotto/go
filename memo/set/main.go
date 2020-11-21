package main

import (
	"fmt"
)

func main() {
	set := make(map[int]struct{})
	set[3] = struct{}{}
	set[2] = struct{}{}

	for k := range set {
		fmt.Println(k)
	}

	_, exist := set[4]
	fmt.Println(exist)
}
