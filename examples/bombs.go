package main

import "fmt"

func printCellState(bomb int) string {
	s := ""
	if 0 < bomb {
		s = fmt.Sprint("%d", bomb)
	} else {
		s = "_"
	}
	return s
}

func bombsPresent(arr []int) bool {
	return getBombCount(arr) > 0
}

func getBombCount(arr []int) int {

	totalBombs := 0
	for n := 0; n < len(arr); n++ {
		totalBombs += arr[n]
	}

	return totalBombs
}

func printNeighborState(matrix []int) {
	bombs := getBombCount(matrix)
	if 0 < bombs {
		fmt.Printf("[%d, _, _, _ ]\n", bombs)
	} else {
		fmt.Println("[_, _, _, _ ]")
	}

}

func printArrayState(matrix []int) {
	fmt.Printf("[")
	for n := 0; n < len(matrix); n++ {
		if n > 0 {
			fmt.Printf(", ")
		}

		fmt.Printf("%d", matrix[n])
	}
	fmt.Printf("]\n")

}

func main2() {

	bombMatrices := [][]int{{0, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}, {0, 0, 0, 1}}

	fmt.Printf("\n\n")
	fmt.Println("beginning game state")
	for _, startArr := range bombMatrices {
		printArrayState(startArr)
	}
	fmt.Printf("\n\n")

	fmt.Printf("\n\n")
	fmt.Println("neighbor states")

	totalArrays := len(bombMatrices)
	for n := 0; n < totalArrays; n++ {
		totalBombs := getBombCount(bombMatrices[n])
		if 0 < totalBombs {
			if n < totalArrays-1 {
				printNeighborState(bombMatrices[n+1])
			} else {
				printNeighborState(bombMatrices[n-1])
			}
		} else {
			printNeighborState([]int{0, 0, 0, 0})
		}
		//printArrayState(bmatrix)
	}

	fmt.Printf("\n\n")
}
