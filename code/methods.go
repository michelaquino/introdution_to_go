package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) ChangeName(newName string) { // método publico
	p.Name = newName
}

func (p Person) doNotChangeName(newName string) { // método privado
	p.Name = newName
}

func main() {
	person := Person{Name: "José"}
	fmt.Println("Idade - 1: ", person.Name)
	person.ChangeName("Maria")
	fmt.Println("Idade - 2: ", person.Name)

	person.doNotChangeName("Antonio")
	fmt.Println("Idade - 3: ", person.Name)
}
