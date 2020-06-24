package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	number bool
	strCount int
	files []string
)

func init() {
	flag.BoolVar(&number, "number", number, "通し番号を付与する")
	flag.IntVar(&strCount, "count", strCount, "各ファイルにおける出力する文字数")
	flag.Parse()
	files = flag.Args()
}

func main() {
	// 実行ファイルのパスを取得
	path, _ := os.Executable()

	//path = filepath.Dir(path)

	for i, file := range files {

		if number {
			fmt.Printf("%v: %v\n", i, file)
		} else {
			fmt.Println(file)
		}
		path, _ := os.Open(filepath.Join(path, file))

		scanner := bufio.NewScanner(path)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

	}
}
