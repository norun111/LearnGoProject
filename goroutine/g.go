package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main")
	go func() {
		fmt.Println("hoge")
	}()
	time.Sleep(time.Second)
}


