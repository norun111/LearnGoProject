package main

import (
	"fmt"
)

func generator(done <-chan interface{}, integers ...int) <-chan int {
	intStream := make(chan int, len(integers))

	go func() {
		defer close(intStream)
		for _, i := range integers {
			select {
			case <-done:
				return
			case intStream <- i:
			}
		}
	}()
	return intStream
}

func multiply(done <-chan interface{}, intStream <-chan int, multiplier int) <-chan int {
	multipleStream := make(chan int)
	go func(){
		defer close(multipleStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case multipleStream <- i*multiplier:
			}
		}
	}()
	return multipleStream
}

func add(done <-chan interface{}, intStream <-chan int, additive int) <-chan int {
	addStream := make(chan int)
	go func() {
		close(addStream)
		for i := range intStream {
			select {
			case <-done:
				return
			case addStream <- i + additive:
			}
		}
	}()
	return addStream
}
func main() {

	done := make(chan interface{})
	defer close(done)
	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}

}
