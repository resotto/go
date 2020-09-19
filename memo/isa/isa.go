package main

import "fmt"

type Person struct {
	Name    string
	Address Address
}

type Address struct {
	Number string
	Street string
	City   string
	State  string
	Zip    string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Location() {
	fmt.Println("I’m at", p.Address.Number, p.Address.Street, p.Address.City, p.Address.State, p.Address.Zip)
}

type Citizen struct {
	Country string
	Person
}

func (c *Citizen) Nationality() {
	fmt.Println(c.Name, "is a citizen of", c.Country)
}

//func SpeakTo(p *Person) {
//    p.Talk() // when citizen given, cannot use &c (type *Citizen) as type *Person in argument to SpeakTo
//}

type Human interface {
	Talk()
}

func SpeakTo(h Human) {
	h.Talk()
}

func main() {
	p := Person{Name: "Dave"}
	c := Citizen{Person: Person{Name: "Steve"}, Country: "America"}

	SpeakTo(&p)
	SpeakTo(&c)
}
