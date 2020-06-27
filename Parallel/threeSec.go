package main

import (
	"fmt"
	"log"
	"sync"
)

func palarell(wg *sync.WaitGroup) {
	defer wg.Done()
	log.Print("sleep1 started.")
	//time.Sleep(1 * time.Second)
	log.Print("sleep1 finished.")
}

func main() {
	wg := new(sync.WaitGroup)

	for i := 0; i < 3; i++{
		wg.Add(1)
		go palarell(wg)
	}

	wg.Wait()
	fmt.Println("All Done!!")
}
