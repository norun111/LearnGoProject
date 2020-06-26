package main

import "fmt"

func main() {
	intStream := make(chan int)

	go func() {
		defer close(intStream)
		intStream <- 1
		for i := 0; i < 10; i++ {
			intStream <- i+1
		}
	}()
	//閉じたチャネルから読み込みは無限にできる
	for i := 0; i < 20; i++ {
		fmt.Println(<-intStream)
	}

	for integer := range intStream {
		fmt.Println(integer)
	}

}
