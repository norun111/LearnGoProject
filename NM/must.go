package main

import (
	"fmt"
	"log"
	"regexp"
)

var validId = regexp.MustCompile(`^[a-z]+\[[1-9]+\]$`)
func main() {
	fmt.Println(validId.MatchString("adam[123]"))

	//関数内で行う場合はエラー処理をする
	validId2, err := regexp.Compile(`^[a-z]+\[[1-9]+\]$`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(validId2.MatchString("adam[333]"))
}