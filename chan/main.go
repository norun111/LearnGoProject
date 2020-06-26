package main

import "fmt"

func goroutine(ch chan string)  {
	ch <- "hello"
}

func main() {
	ch := make(chan string)
	go goroutine(ch)

	fmt.Println(<-ch)
}
