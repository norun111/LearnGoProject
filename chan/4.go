package main

import "fmt"

func main() {
	chanOwner := func() <-chan int {
		intStream := make(chan int)

			go func() {
				defer close(intStream)
				for i := 0; i < 5; i++ {
					intStream <- i
				}
			}()

		return intStream
	}

	streamList := chanOwner()
	for i := range streamList {
		fmt.Println(i)
	}
	fmt.Println("done")
}
