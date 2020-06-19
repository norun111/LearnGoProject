package main

import (
	"fmt"
	"time"
)

func main() {
	//バッファがないとランタイムパニックが発生
	//最後にerrorが出力されているのは、closeされているから
	ch := make(chan int, 3)

	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
	}()

	time.Sleep(time.Second)
	close(ch)

	for {
		a, ok := <-ch
		if !ok {
			fmt.Println("error")
			break
		}
		fmt.Println(a)
	}
}
