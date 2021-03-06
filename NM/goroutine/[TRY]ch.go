package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func makeCh() chan string {
	return make(chan string)
}

func input(r io.Reader) <-chan string { // TODO: チャネルを作る
	ch := makeCh()
	go func() {
		s := bufio.NewScanner(r)
		for s.Scan() {
			// TODO: チャネルに読み込んだ文字列を送る
			ch <- s.Text()
		}
			// TODO: チャネルを閉じる
			close(ch)
	}()
			// TODO: チャネルを返す
			return ch
}

func main() {
	ch := input(os.Stdin)
	for {
		fmt.Print(">")
		fmt.Println(<-ch)
	}
}
