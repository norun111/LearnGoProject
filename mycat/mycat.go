package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var n = flag.Bool("n", false, "通し番号を付与する")
	flag.Parse()

	var (
		files = flag.Args()
		path, _ = os.Executable()
	)

	path = filepath.Dir(path)

	i := 1
	for x := 0; x < len(files); x++ {
		sf, _ := os.Open(filepath.Join(path, files[x]))
		scanner := bufio.NewScanner(sf)

		for ;scanner.Scan(); i++ {
			if *n{
				//オプションがある場合
				fmt.Printf("%v:\n", i)
			}
			fmt.Println("aa")
			fmt.Println(scanner.Text())
		}
	}
}