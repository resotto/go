package main

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	s := rand.Intn(6) + 1
	// println(s)

	switch {
	case s == 6:
		println("大吉")
	case s == 5 || s == 4:
		println("中吉")
	case s == 3 || s == 2:
		println("吉")
	case s == 1:
		println("凶")
	}
}
