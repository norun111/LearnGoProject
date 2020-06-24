package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := filepath.Join("dir", "main.go")
	//拡張子取得
	fmt.Println(filepath.Ext(path))
	//ファイル名取得
	fmt.Println(filepath.Base(path))
	//ディレクトリ名取得
	fmt.Println(filepath.Dir(path))
}
