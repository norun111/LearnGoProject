package main

import (
	"fmt"
	"sync"
)
//sayHello関数をホストしているゴルーチンが終了するまでメインゴルーチンをブロックする
func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait() //合流ポイント
}
