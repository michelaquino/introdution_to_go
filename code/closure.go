package main

import "fmt"

func main() {
	sub := func(x, y int) int {
		return x - y
	}
	fmt.Println(soma(10, 2))
}
