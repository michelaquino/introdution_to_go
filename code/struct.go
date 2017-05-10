package main

import "fmt"

type Vehicle struct {
	Name   string
	Engine Engine
	Age    int
}

type Engine struct {
	Maker string
}

func main() {
	vehicle := Vehicle{
		Name:   "Fusca",
		Engine: Engine{Maker: "Ferrari"},
		Age:    40,
	}

	fmt.Println("Nome do veiculo", vehicle.Name)
	fmt.Println("Idade", vehicle.Age)
	fmt.Println("Fabricante do motor", vehicle.Engine.Maker)
}
