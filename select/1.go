package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan interface{})
	close(c1) //初期化されたチャネルをcloseで開放している
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int

	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		case <-time.After(1 * time.Second):
			fmt.Println("Time out")
		}
	}
	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}
