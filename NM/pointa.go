package main

import "fmt"

func f(xp *int) {
	*xp = 100
}
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	var x int
	f(&x)
	println(x)

	n, m := 10, 20
	swap(&n, &m)
	fmt.Println(n, m)
}
