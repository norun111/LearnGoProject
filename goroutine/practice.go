package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start!")
	ch := make(chan bool)

	go func() {
		defer close(ch)
		time.Sleep(2*time.Second)
		ch <- true
	}()

	fmt.Println(<-ch)
	fmt.Println("End!")
}
