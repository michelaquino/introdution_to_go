package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for {
		fmt.Println("Executando task ", name)
		time.Sleep(time.Second)
	}
}

func main() {
	go task("TASK 1")
	go task("TASK 2")
	go task("TASK 3")

	task("main")
}
