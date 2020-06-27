package main

import "fmt"

func main() {
	multiple := func(values []int, multiplier int) []int {
		multipleSlice := make([]int, len(values))

		for i, v := range values {
			multipleSlice[i] = v * multiplier
		}
		return multipleSlice
	}

	additive := func(values []int, additive int) []int {
		additiveSlice := make([]int, len(values))

		for i, v := range values {
			additiveSlice[i] = v + additive
		}
		return additiveSlice
	}

	ints := []int{1, 2, 3, 4}
	//multipleSliceを元にadditive関数を実行
	for _, v := range additive(multiple(ints, 2), 1){
		fmt.Println(v)
	}
}

