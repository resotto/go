package main

type A struct {
}

type B struct {
	A //B is-a A
}

func save(A) {
	//do something
}

func main() {
	b := B{}
	save(&b) // cannot use &b (type *B) as type A in argument to save
}
