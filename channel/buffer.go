package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "Hello world"
}

func callerB(c chan string) {
	c <- "Hola Mundo"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)

	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)
		select {
		case msg := <-a:
			fmt.Println(msg)
		case msg := <-b:
			fmt.Println(msg)
		default:
			fmt.Println("error")
		}
	}
}
