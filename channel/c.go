package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int)
	//go func() {
	//	ch <- 1
	//}()
	//a := <-ch
	//fmt.Println(a)

	ch := make(chan int, 5)

	go func(){
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	time.Sleep(time.Second)

	for i := 1; i <= 5; i++ {
		tmp := <-ch

		fmt.Println(tmp)
	}

	fmt.Println("witing")
	time.Sleep(time.Second)

	for i := 1; i <= 5; i++ {
		tmp := <-ch
		fmt.Println(tmp)
	}


}