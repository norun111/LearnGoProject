package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func main() {
	result := testing.Benchmark(func(b *testing.B) {run("A", "B", "C", "D", "E")})
	fmt.Println(result.T)
}

func run(name ...string) {
	fmt.Println("Start!")
	wg := new(sync.WaitGroup)

	isFin := make(chan bool, len(name))

	for _, v := range name {
		wg.Add(1)
		go process(v, isFin, wg)
	}

	wg.Wait()
	close(isFin)
	fmt.Println("Finish!")
}

func process(name string, isFin chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2*time.Second)
	fmt.Println(name)
	isFin <- true
}
