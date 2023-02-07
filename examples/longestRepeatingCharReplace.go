package main

import (
	"fmt"
)

/*
Given a string s and an integer k. You can choose any character of the string and change it to
any other uppercase English character. You can perform this operation at most k times.

Return the length of the longest substring containing the same letter you can get after performing
the above operations.

Input: s = "ABAB", k = 2
Output: 4
Explanation: Replace the two 'A's with two 'B's or vice versa.

Input: s = "AABABBA", k = 1
Output: 4
Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
The substring "BBBB" has the longest repeating letters, which is 4.

Constraints
N ~ length of input string
k ~ Number of operations

Time Complexity: O(N)
Auxiliary Space Complexity: O(k)
*/

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func longestRepeatingCharReplace(s string, k int) int {

	max_value := 0
	l := 0
	result := 0
	characters := make(map[string]int)

	for r, c := range s {

		characters[string(c)]++
		max_value = Max(max_value, characters[string(c)])
		for r-l+1-max_value > k {
			characters[string(s[l])]--
			l++
		}
		result = Max(result, r-l+1)
	}
	return result
}

func main() {
	/*
		A  A  B  A  B  B  A
		l0
		r0
		k1

		Input: ABAB, k = 2      =>	Output: 4
		Input: AABABBA, k = 1	 	=>	Output: 4
		Input: AABABBA, k = 3	 	=>	Output: 7
	*/
	fmt.Println(longestRepeatingCharReplace("ABAB", 2))
	fmt.Println(longestRepeatingCharReplace("AABABBA", 1))
	fmt.Println(longestRepeatingCharReplace("AABABBA", 3))
}
