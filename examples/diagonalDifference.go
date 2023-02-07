package main

import (
	"fmt"
	"math"
)

func diagonalDifference(arr [][]int32) int32 {
	ltr := 0
	rtr := 0

	for i := 0; i < len(arr); i++ {
		ltr += int(arr[i][i])
		fmt.Printf("%d ", arr[i][i])
	}
	fmt.Println()

	for i, j := 0, len(arr)-1; j >= 0 && i < len(arr); j-- {
		rtr += int(arr[i][j])
		fmt.Printf("%d ", arr[i][j])
		i += 1
	}

	fmt.Println()
	return int32(math.Abs(float64(ltr - rtr)))
}

func main() {

	fmt.Printf("%d", diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 9}}))
}
