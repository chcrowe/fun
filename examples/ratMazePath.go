/*
Given a maze, represented by a matrix of bits (values 0 and 1), a rat must find a path from index [0][0] to [n-1][m-1]. The rat can only travel to the right or down, and can only travel on 0 values. 1 values represent walls.

Input:   [[0, 0, 0, 1],

	[0, 1, 0, 1],
	[0, 1, 0, 0],
	[0, 0, 1, 0]]

Output:

	[[0, 0],
	[0, 1],
	[0, 2],
	[1, 2],
	[2, 2],
	[2, 3],
	[3, 3]]

For M x N matrix.
Time Complexity: 0(MN)
Auxiliary Space Complexity: O(MN)

base cases:

	check if i,j is not in bounds: return false
	check if i,j is a 1: return false
	check if i,j is lower right corner:
	    print out the list and stop (or return true)

recursive case:

	call myself with row+1, col, append [row+1,col] to list
	if that fails (the call returned, or returned false)
	   remove [row+1,col] from list
	   call myself with row, col+1 appended to list

0,0, [[0,0]]

	1,0, [[0,0], [1,0]]
	        push [2,0]
	   2,0, [[0,0], [1,0], [2,0]]  ^^
	       pop the [2,0]
	       push [1,1]
	   1,1, [[0,0], [1,0], [1,1]]
	      ...
	         ...
	              [3,3]

PSeudocode
iterate the matrix from top left to bottom right
store each test coordinate in path

while matrix row and col are within limits

	we will test for a zero
	if 0
	  valid move (adding the coordinate to output)
	else
	  we increment right 1 or down 1 if out of bounds
*/
package main

import (
	"fmt"
)

func findPath(matrix [][]int, path []string, row int, col int) {

	//if out of bounds or cell contains a 1 return
	if row >= len(matrix) || col >= len(matrix[0]) {
		return
	} else if 1 == matrix[row][col] {
		//fmt.Printf("%v (%d) %v\n", []int{row, col}, matrix[row][col], path)
		return
	}

	//fmt.Printf("%v (%d)\n", []int{row, col}, matrix[row][col])
	//path.Push([]int{row, col})
	path = append(path, fmt.Sprintf("%v", []int{row, col}))

	//if end of matrix bounds, exit
	if row == len(matrix)-1 && col == len(matrix[0])-1 {
		//fmt.Printf("%v\n", path)
		result = path[0:]
		return
	}

	//recurse down
	findPath(matrix, path, row+1, col)
	//recurse right
	findPath(matrix, path, row, col+1)

	//path.Pop()
	path = path[:len(path)-1]

	//uncomment to make faster
	//matrix[row][col] = 1
}

var result []string

func main() {

	maze := [][]int{{0, 0, 0, 1}, {0, 1, 0, 1}, {0, 1, 0, 0}, {0, 0, 1, 0}}

	for _, row := range maze {
		fmt.Printf("%v\n", row)
	}

	fmt.Println()
	//route := stack.New()
	route := make([]string, 0)
	findPath(maze, route, 0, 0)

	fmt.Println()
	fmt.Printf("%v\n", result)
	fmt.Println()
	for i, row := range maze {
		fmt.Printf("%d %v\n", i, row)
	}
	fmt.Println()
	fmt.Println()

	//fmt.Printf("route...\n%v\n", route)

}
