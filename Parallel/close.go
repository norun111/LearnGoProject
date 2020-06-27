package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	done := make(chan struct{})

	for i:=0; i<10; i++{
		wg.Add(1)
		go func(i int) {
			<- done
			fmt.Println("done", i)
			defer wg.Done()
		}(i)
	}

	close(done)
	wg.Wait()
}
