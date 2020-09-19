package main

import "fmt"

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	return 0, nil
}

// ErrNegativeSqrt is type
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g\n", float64(e))
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
