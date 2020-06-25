package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var m sync.Mutex

	increment := func() {
		m.Lock()
		count++
		defer m.Unlock() //Lockが保持しているクリティカルセクションでの処理が終了

		fmt.Println(count)
	}

	decrement := func() {
		m.Lock()
		count--
		defer m.Unlock()
		fmt.Println(count)
	}

	var arithmetic sync.WaitGroup

	//インクリメント
	for i := 0; i <= 5; i++{
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	//デクリメント
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	fmt.Println("arithmetic completed")
}
