package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Helle\n")

	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	read1, err := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	file, _ := os.Create("data2")
	defer file.Close()


}
