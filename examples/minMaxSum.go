package main

import "fmt"

func miniMaxSum(arr []int32) {
	// Write your code here

	var total int64 = 0
	for _, n := range arr {
		total += int64(n)
	}

	var adjTotals []int64
	var tot int64 = 0
	var min int64 = 0
	var max int64 = 0

	for _, n := range arr {
		tot = total - int64(n)
		if tot > max {
			max = tot
		}
		if 0 == min || tot < min {
			min = tot
		}

		//fmt.Printf("%d ", tot)
		adjTotals = append(adjTotals, total-tot)
	}

	fmt.Printf("\n%d %d", min, max)

}

func main() {

	miniMaxSum([]int32{1, 2, 3, 4, 5})
}
