package main

import "fmt"

//int型を新たにHex型として定義している
type Hex int

type Stringer interface {
	String() string
}


func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	var s Stringer = Hex(100)
	fmt.Println(s)
}
