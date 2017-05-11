package main

import (
	"fmt"
	"time"
)

func task(name string) {
	i := 0
	for {
		fmt.Println("Executando task ", name, " - ", i)
		time.Sleep(time.Second)
		i++
	}
}

func main() {
	go task("TASK 1")
	go task("TASK 2")
	go task("TASK 3")

	task("main")
}
