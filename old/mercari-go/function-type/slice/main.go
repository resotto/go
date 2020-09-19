package main

import "fmt"

func main() {
	// ns := make([]int, 3, 10)
	// println(ns[2])
	// println(cap(ns))
	// println(len(ns))
	// ns = append(ns, 60, 70)
	// println(ns[3])
	// println(cap(ns))

	// a := []int{10, 20}
	// fmt.Println(a, cap(a))

	// b := append(a, 30)
	// a[0] = 100
	// fmt.Println(b, cap(b))

	// c := append(b, 40)
	// b[1] = 200
	// fmt.Println(c, cap(c))

	// s := []int{10, 20, 30, 40, 50}
	// n, m := 2, 4

	// fmt.Println(s[n:])
	// fmt.Println(s[:m])
	// ms := s[:m:m]
	// fmt.Println(ms, cap(ms))
	slice := []int{19, 86, 1, 12}
	var sum int
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	fmt.Println(sum)
}
