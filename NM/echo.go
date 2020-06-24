package main

import (
	"fmt"
	"os"
)

func main() {
	buff := make([]byte, 256)
	for {
		c, _ := os.Stdin.Read(buff)
		fmt.Println(c)
		if c == 0 { break }
		os.Stdout.Write(buff[:c])
	}
}