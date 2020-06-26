package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(5*time.Second)
		close(c) //5秒待った後にチャネルを閉じます
	}()

	fmt.Println("Blocking")
	select {
	case <-c:
		fmt.Printf("Unblocking %v\n", time.Since(start))
	}
}
