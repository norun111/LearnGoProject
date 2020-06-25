package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int

	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated++
			mem := make([]byte, 1024)
			return &mem
		},
	}
	//プールに4つ確保する
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024

	fmt.Println(numWorkers)

	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := numWorkers; i > 0; i-- {
		go func(i int) {
			defer wg.Done()
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
		}(i)
	}

	wg.Wait()
	fmt.Printf("%d calculate were created\n", numCalcsCreated)
}
