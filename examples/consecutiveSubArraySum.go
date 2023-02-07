package main

import (
	"fmt"
)

/*
Given an array of positive integers and a target value, return true if there is a subarray of consecutive
elements that sum up to this target value.
Example
Input: [6,12,1,7,5,2,3], 14      	=>	Output: true (7+5+2)
Input: [8,3,7,9,10,1,13], 50		=>	Output: false

Constraints
Time Complexity: O(N)
Auxiliary Space Complexity: O(1)
All elements are positive
*/
// alternate approach
// func consecutiveSubArraySum(nums []int, k int) bool {

// 	sum := k
// 	l := 0

// 	for r := 0; r < len(nums); r++ {
// 		sum -= nums[r]

// 		if 0 == sum {
// 			return true
// 		} else if 0 > sum {
// 			l++
// 			r = l - 1
// 			sum = k
// 		}
// 	}

// 	return false
// }

func consecutiveSubArraySum(nums []int, k int) bool {

	sum := 0
	hash := make(map[int]int)

	for r, val := range nums {
		sum += val

		if 0 != k {
			sum %= k
		}
		_, ok := hash[sum]
		if true == ok {
			if r-hash[sum] > 1 {
				return true
			}
		} else {
			hash[sum] = r
		}
	}

	return false
}

func main() {
	/*
		Input: [6,12,1,7,5,2,3], 14      	=>	Output: true (7+5+2)
		Input: [8,3,7,9,10,1,13], 50		=>	Output: false
	*/
	fmt.Println(consecutiveSubArraySum([]int{6, 12, 1, 7, 5, 2, 3}, 14))
	//	fmt.Println(consecutiveSubArraySum([]int{8, 3, 7, 9, 10, 1, 13}, 50))
}
