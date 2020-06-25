package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		ch <- 100
	}()

	go func() {
		v := <-ch
		fmt.Println(v)
	}()

	time.Sleep(2 * time.Second)
}
