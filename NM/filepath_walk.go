package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func Walk(path string, info os.FileInfo, err error) error {
	err = filepath.Walk("dir", func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ".go" {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {

}
