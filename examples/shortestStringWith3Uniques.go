package main

import (
	"fmt"
	"math"
)

/*
Given a string s, find the longest substring without repeating characters.

Input: String
Output: String

Input: abcabcbb      =>	Output: abc
Input: bbbbb	 	=>	Output: b
Input: pwwkew		=>	Output: wke

01 loop through each individual character of the main string
02 create working string that begins at that index up to the remainder of string
03 use a map to track the number of times an individual char occurs within the string
05 once we have identified each of the search characters
06 do bookkeeping to record the shortest string that contains all of the characters

a b c a b c b b
^

	^
*/

func shortestStringWith3Uniques(s string) string {

	l := 0
	best_l, best_r := 0, math.MaxInt

	uniques := 0
	characters := make(map[string]int)

	for r, c := range s {

		//hunt phase
		if uniques < 3 {
			_, ok := characters[string(c)]
			if false == ok {
				uniques++
			}
			characters[string(c)]++
		}

		//catchup phase
		for 3 <= uniques {
			if r-l < best_r-best_l {
				best_l, best_r = l, r
			}

			if 1 == characters[string(s[l])] {
				uniques--
			}
			characters[string(s[l])]--
			l++
		}
	}

	if math.MaxInt != best_r {
		return s[best_l : best_r+1]
	} else {
		return ""
	}
}

func main() {
	/*
	   Input: aabaca => Output: bac
	   Input: abaacc => Output: baac
	   Input: abc => Output: abc
	   Input: aabb => Output: False
	*/
	fmt.Println(shortestStringWith3Uniques("aabaca"))
	fmt.Println(shortestStringWith3Uniques("abaacc"))
	fmt.Println(shortestStringWith3Uniques("abc"))
	fmt.Println(shortestStringWith3Uniques("aabb"))
}
