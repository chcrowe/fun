package main

import (
	"bufio"
	"fmt"
	"os"
)

func PrintBoard(bd [][]string) {

	for _, row := range bd {
		for _, col := range row {
			fmt.Printf(" %-1s ", col)
		}
		fmt.Println()
	}
}

func checkForWinner(bd [][]string, mark string) (bool, string) {
	/*
		search for 3 X's in any row
		search for 3 X's in any col
		search for 3Xs from 0,0 to 2,2
		search for 3Xs from 2,0 to 0,2
		if any match, exit true	w coords
	*/
	const maxLen int = 3

	//search for 3 X's in any row
	runningCount := 0
	for r, row := range bd {
		runningCount = 0
		for c, col := range row {
			if mark == col {
				runningCount++
			} else {
				break
			}
			if maxLen == runningCount {
				return true, fmt.Sprintf("%q wins... 3 in a row (%d,0 - %d,%d)", mark, r, r, c)
			}
		}
	}

	//search for 3 X's in any column
	for col := 0; col < maxLen; col++ {
		runningCount = 0
		for row := 0; row < maxLen; row++ {
			if mark == bd[row][col] {
				runningCount++
			} else {
				break
			}
			if maxLen == runningCount {
				return true, fmt.Sprintf("%q wins... 3 in a column (0,%d - %d,%d)", mark, col, row, col)
			}
		}
	}

	runningCount = 0
	for n := 0; n < maxLen; n++ {
		if mark == bd[n][n] {
			runningCount++
		} else {
			break
		}
		if maxLen == runningCount {
			return true, fmt.Sprintf("%q wins... 3 in diagonal ([0,0][1,1][2,2])", mark)
		}
	}

	runningCount = 0
	for r, c := 0, 2; r < maxLen; r++ {
		if mark == bd[r][c] {
			runningCount++
		} else {
			break
		}
		if maxLen == runningCount {
			return true, fmt.Sprintf("%q wins... 3 in diagonal ([0,2][1,1][2,0])", mark)
		}

		c--
	}

	return false, ""
}

func SpotAvailable(bd [][]string, row int, col int) bool {
	return bd[row][col] == "-"
}

func MoveWithinBounds(bd [][]string, row int, col int) bool {
	maxRows := len(bd)
	maxCols := len(bd[0])

	return 0 <= row && row < maxRows && 0 <= col && col < maxCols
}

func GetNextMove(move int) (error, string, int, int) {

	playerNo := 1
	mark := "O"
	rem := move % 2
	if 0 == rem {
		mark = "X"
	} else {
		playerNo = 2
	}

	row, col := 0, 0
	fmt.Printf("\nPlayer %d-%q move (row:col): ", playerNo, mark)
	line, err := reader.ReadString('\n')
	if err == nil && 3 <= len(line) {
		fmt.Sscanf(line, "%01d %01d", &row, &col)
		//fmt.Printf("%03d-%03d\n", row, col)
	} else {
		fmt.Println("# input error... try again")
	}

	return err, mark, row, col
}

var board [][]string
var reader *bufio.Reader

func main() {
	board = [][]string{{"-", "-", "-"}, {"-", "-", "-"}, {"-", "-", "-"}}

	reader = bufio.NewReader(os.Stdin)
	PrintBoard(board)

	move := 0
	for {
		err, mark, row, col := GetNextMove(move)
		if nil == err && MoveWithinBounds(board, row, col) {
			if false == SpotAvailable(board, row, col) {
				fmt.Println("Spot occupied... choose again")
			} else {
				board[row][col] = mark
				move++
				PrintBoard(board)
				res, msg := checkForWinner(board, mark)
				if res {
					fmt.Println("*******************************************")
					fmt.Println(msg)
					fmt.Println("*******************************************")
					break
				}
			}
		}
	}
}
