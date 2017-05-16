package main

import "fmt"

func main() {
	people := map[string]int{
		"Jose":  25,
		"Maria": 26,
	}

	people["Jose"] = 10

	for name, age := range people {
		fmt.Printf("Chave %s do mapa valor: %d\n", name, age)
	}

	delete(people, "Jose")
	fmt.Printf("Posição não alocada: %d\n", people["Jose"])
}
