package main

import "sync"

func main() {
	pool := &sync.Pool{
		//全てやりとりはinterface型で行う
		New: func() interface{} {
			return 1
		},
	}
	//型アサーションが必要
	r, ok := pool.Get().(int)
	if !ok {
		panic("something wrong")
	}
	pool.Put(r)
}
