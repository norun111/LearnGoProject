package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	//sc := bufio.NewScanner(os.Stdin)
	//// 1行ずつ読み込んで繰り返す
	//for sc.Scan() {
	//	//1行分を出力する
	//	fmt.Println(sc.Text())
	//}
	//
	//if err := sc.Err(); err != nil {
	//	fmt.Println(os.Stderr, "loading err:", err)
	//}

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for {
		c, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		w.WriteByte(c)
		if c == '\n' { w.Flush() }
	}
	w.Flush()
}
