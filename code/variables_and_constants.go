package main

import "fmt"

const CONSTANT = 1

var (
	a_out string
	b_out int64 = 2
)

func main() {
	var x int // ou var x int = 10
	x = 10
	fmt.Printf("O valor de 'x' é %d\n", x)

	a := 20
	fmt.Printf("O valor de 'a' era %d\n", a)

	fmt.Println("O valor de 'a_out' é: ", a_out)
	fmt.Println("O valor de 'b_out' é: ", b_out)

	fmt.Println("Imprimindo constante: ", CONSTANT)
}
