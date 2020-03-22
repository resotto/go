package main

func main() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			println(i, "- 偶数")
		} else {
			println(i, "- 奇数")
		}
	}

	println("---")

	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0:
			println(i, "- 偶数")
		case i%2 == 1:
			println(i, "- 奇数")
		default:
			println("defualt")
		}
	}
}
