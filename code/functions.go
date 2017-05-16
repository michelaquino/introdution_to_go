package main

import "fmt"

// exemplo de função privada
func add(x int, y int) int {
	return x + y
}

// exemplo de função que retorna mais que um valor
func Invert(x, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("O resultado de 10 + 20 é", add(10, 20))
	a, b := Invert(10, 20)
	fmt.Println("O resultado de Invert(10, 20) é", a, b)
}
