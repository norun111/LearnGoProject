package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("hoge.txt")
	p := make([]byte, 16)

	n, err := f.Read(p)
	fmt.Println(*f)
	fmt.Println(n)

	if 16 < n {
		log.Fatalf("%dバイト読もうとしましたが、%dバイトしか読めませんでした\n", n, 16)
	}
	if err != nil {
		log.Fatalf("読込中にエラーが発生しました：%v\n", err)
	}
}
