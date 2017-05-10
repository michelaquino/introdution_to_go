package main

import "fmt"

func main() {
	defer func() {
		if str := recover(); str != nil {
			fmt.Println("Recovered: ", str)
		}
	}()

	panic("BREAKED!")
}
