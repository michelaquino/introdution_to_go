package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("DEFER")
	}()

	fmt.Println("ULTIMA LINHA")
}
