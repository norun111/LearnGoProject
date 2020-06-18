package main

import (
	"flag"
	"fmt"
)

func main() {
	//String
	//var name = flag.String("name", "Tomoya", "please")
	//flag.Parse()
	//fmt.Println("The name flag is", *name)

	//Int
	var num int
	flag.IntVar(&num, "num", 100, "Please")
	flag.Parse()
	fmt.Println("The num is", num)
}
