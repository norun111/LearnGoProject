package main

import (
	"flag"
	"fmt"
	"strings"
)

var msg = flag.String("msg", "default", "explanation")
var n int

func init() {
	flag.IntVar(&n, "n",1, "count")
}

func main() {
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))
}
