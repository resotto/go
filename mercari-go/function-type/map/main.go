package main

func main() {
	m := map[string]int{"x": 10, "y": 20}
	println(m["x"])
	m["z"] = 30
	n, ok := m["z"]
	println(n, ok)
	delete(m, "z")
	n, ok = m["z"]
	println(n, ok)
}
