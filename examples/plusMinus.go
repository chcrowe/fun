package main

import "fmt"

func plusMinus(arr []int) {
	// Write your code here
	pos := 0
	neg := 0
	zero := 0

	for _, n := range arr {
		fmt.Printf("%d, ", n)
		if 0 < n {
			pos++
		} else if 0 > n {
			neg++
		} else {
			zero++
		}
	}
	fmt.Printf("\npos %d, neg %d, zero %d\n\n", pos, neg, zero)

	results := []int{pos, neg, zero}
	nums := len(arr)

	for _, j := range results {
		fmt.Printf("%.6f\n", float32(j)/float32(nums))
	}

}

func main() {

	plusMinus([]int{-4, 3, -9, 0, 4, 1})
}
