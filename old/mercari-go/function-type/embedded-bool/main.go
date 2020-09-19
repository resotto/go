package main

func main() {
	var a, b, c bool
	// a, b = true, true
	// c = every

	// a, b = every patterns
	// c = false

	if a && b || !c {
		println("true")
	} else {
		println("false")
	}
}
