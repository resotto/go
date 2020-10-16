package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

type Citizen struct {
	Person
	location string
}

func main() {
	c := Citizen{Person{"John", 32}, "NY"}
	//	c := Citizen{"NY", Person{"John", 32}} // error, argument order invalid
	//	c.name = "John"
	//	c.age = 32
	fmt.Println(c.name, c.age, c.location)               // John 32 NY
	fmt.Println(c.Person.name, c.Person.age, c.location) // John 32 NY
}
