package main

import (
	"flag"
	"fmt"
)

func main(){
	//flag.Parse()
	//args := flag.Args()
	//fmt.Println(args)

	var (
		i = flag.Int("int", 0, "int flag")
		s = flag.String("str", "default", "string flag")
		b = flag.Bool("bool", false, "bool flag")
	)

	flag.Parse()
	fmt.Println(*i, *s, *b)
}
