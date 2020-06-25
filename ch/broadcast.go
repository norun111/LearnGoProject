package main

import (
	"fmt"
	"sync"
)

type Button struct {
	Clicked *sync.Cond
}

func main() {
	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, fn func()){
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)

		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			c.Wait()
			defer c.L.Unlock()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegister sync.WaitGroup
	clickRegister.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Max")
		clickRegister.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Display")
		clickRegister.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse")
		clickRegister.Done()
	})

	button.Clicked.Broadcast()

	clickRegister.Wait()
}
