package main

import "fmt"

func main() {
	x := 5
	zero(&x)
	fmt.Println("x: ", x)

	y := new(int)
	fmt.Println("y - endereço de memoria: ", y)
	fmt.Println("y - valor armazenado: ", *y)

	*y = 1
	fmt.Println("y - endereço de memoria: ", y)
	fmt.Println("y - valor armazenado: ", *y)
}

func zero(x *int) {
	*x = 0
}
