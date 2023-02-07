/**
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:


Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
Example 2:

Input: numRows = 1
Output: [[1]]

1
1 1
1 2 1
1 3 3 1
-------
0 1 2 3

        1          1
      1   1        2
    1   2   1      3
  1   3   3   1    4 row index 2 = prevrow[1-0] + prevrow[2-1]

  create an matrix of n arrays
  build rows until n rows
  if row items gtr than 2, add parent node to calculate next item

*/

package main

import (
	"fmt"
)

func main() {

	totRows := 7
	m := make([][]int, totRows)

	for i := 0; i < totRows; i++ {
		//dimension the next row
		m[i] = make([]int, i+1)
		for j := 0; j < len(m[i]); j++ {
			if 0 == j || j == len(m[i])-1 {
				m[i][j] = 1
			} else if 0 < j {
				m[i][j] = m[i-1][j-1] + m[i-1][j]
			}
		}
		for r := 0; r < totRows-i; r++ {
			fmt.Printf("%s", " ")
		}
		fmt.Printf("%v\n", m[i])
	}

}

/*
         [1]
        [1 1]
       [1 2 1]
      [1 3 3 1]
     [1 4 6 4 1]
    [1 5 10 10 5 1]
   [1 6 15 20 15 6 1]
  [1 7 21 35 35 21 7 1]
[1 8 28 56 70 56 28 8 1]
*/
