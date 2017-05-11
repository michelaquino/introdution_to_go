package main

import (
	"fmt"
	"time"
)

func task(name string, channel chan int) {
	for {
		valuePassed := <-channel
		fmt.Println("Executando task ", name, " com valor ", valuePassed)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)
	go task("TASK 1", channel)
	go task("TASK 2", channel)

	for i := 0; i < 10; i++ {
		channel <- i
	}
}
