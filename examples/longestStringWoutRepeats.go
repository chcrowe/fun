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

func longestStringWoutRepeats(s string) string {

	l := 0
	best_l, best_r := 0, 0

	repeated := 0
	characters := make(map[string]int)

	for r, c := range s {

		//expansion phase
		if repeated == 0 {
			if r-l > best_r-best_l {
				best_l, best_r = l, r
			}
			characters[string(c)]++
			if 1 < characters[string(c)] {
				repeated++
			}
		}

		//contraction phase
		for 0 < repeated {

			characters[string(s[l])]--
			if 1 == characters[string(s[l])] {
				repeated--
			}
			l++
		}
	}

	if math.MaxInt != best_r {
		return s[best_l:best_r]
	} else {
		return ""
	}
}

func main() {
	/*
		Input: abcabcbb      =>	Output: abc
		Input: bbbbb	 	=>	Output: b
		Input: pwwkew		=>	Output: wke
	*/
	fmt.Println(longestStringWoutRepeats("abcabcbb"))
	fmt.Println(longestStringWoutRepeats("bbbbb"))
	fmt.Println(longestStringWoutRepeats("pwwkew"))
}
