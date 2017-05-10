package main

import "fmt"

type Animal interface {
	Shout() string
}

func callMyAnimal(a Animal) {
	fmt.Println("My animal says:", a.Shout())
}

type Duck struct{}

func (c Duck) Shout() string { return "quá" }
func (c Duck) Walk()         {}
func (c Duck) Eat()          {}

func main() {
	var animal Animal

	animal = Duck{}
	callMyAnimal(animal)
}
