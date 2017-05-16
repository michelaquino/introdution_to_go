package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("DEFER")
		fmt.Println("DEFER 1")
	}()

	fmt.Println("ULTIMA LINHA")
}
