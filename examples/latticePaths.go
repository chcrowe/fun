package main

import (
	"fmt"
)

func latticePaths(row int32, col int32) int32 {
	if row < 0 || col < 0 {
		return 0
	}
	if row == 0 && col == 0 {
		return 1
	}

	up := latticePaths(row-1, col)
	left := latticePaths(row, col-1)

	return up + left
}

func main() {
	fmt.Println(latticePaths(2, 3))
}
