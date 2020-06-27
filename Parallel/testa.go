package main

import (
	"fmt"
	"time"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go testa()

	wg.Wait()
}

func testa() {
	for {
		fmt.Println("test goroutine")
		time.Sleep(2 * time.Second)
	}
}
