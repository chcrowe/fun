package main

import "fmt"

func countingSort(arr []int32) []int32 {
	a := make([]int32, 100)

	for _, n := range arr {
		a[n]++
	}

	return a
}

func main() {
	fmt.Println("a: ", countingSort([]int32{1, 1, 3, 2, 1}))
}
