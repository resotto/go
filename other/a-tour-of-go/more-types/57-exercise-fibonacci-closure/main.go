package main

import "fmt"

// fibonacci is a function that returns
// a function that returns int.
func fibonacci() func() int {
	i := 0
	seq := make([]int, 0)
	return func() int {
		if i == 0 {
			seq = append(seq, 0)
			i++
			return seq[0]
		} else if i == 1 {
			seq = append(seq, 1)
			i++
			return seq[0] + seq[1]
		}
		current := seq[i-2] + seq[i-1]
		seq = append(seq, current)
		i++
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
