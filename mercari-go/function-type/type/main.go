package main

func main() {
	type MyInt int
	var n int = 100
	m := MyInt(n)
	n = int(m)
	println(n)

	type Score struct {
		UserID string
		GameID int
		point  int
	}
}
