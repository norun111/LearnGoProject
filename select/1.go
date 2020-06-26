package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c1 := make(chan interface{})
	close(c1) //初期化されたチャネルをcloseで開放している
	c2 := make(chan interface{})
	close(c2)

	//var c3 chan<- int //panic: close of nil channel
	//close(c3)

	var c1Count, c2Count int

	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		case <-time.After(1 * time.Second):
			fmt.Println("Time out")
		default:
			fmt.Printf("\nIn default after %v\n", time.Since(start))
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}
