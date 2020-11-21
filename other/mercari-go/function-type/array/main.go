package main

func main() {
	var ns1 [5]int
	var ns2 = [5]int{10, 20, 30, 40, 50}
	ns3 := [...]int{10, 20, 30, 40, 50}
	ns4 := [...]int{5: 50, 10: 100}

	println(len(ns1))
	println(ns2[2])
	println(ns3[3])
	println(ns4[5])
	println(ns4[1:6])
}
