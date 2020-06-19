package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		ch1 <- "cat"
	}()

	go func() {
		ch2 <- "dog"
	}()

	time.Sleep(time.Second)

	for i := 0; i < 2; i++ {
		select {
		case a := <-ch1:
			fmt.Println(a)
		case b := <-ch2:
			fmt.Println(b)
			default:
				fmt.Println("nothing")

		}
	}
}
