package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)
func main() {
	fmt.Println("send me Signal")

	l := sync.Mutex{}
	//&lを引数とする
	c := sync.NewCond(&l)

	for i:=0; i<10; i++ {
		go func(i int) {
			l.Lock()
			defer l.Unlock()
			c.Wait()
			fmt.Printf("student%d「Here！」\n", i)
		}(i)
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	if scanner.Text() == "Are you here?" {
		c.Broadcast()
	}
	time.Sleep(1*time.Second)
}
