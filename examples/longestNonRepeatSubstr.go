package main

import (
	"fmt"
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
func longestNonRepeatSubstr(fullString string) string {
	maxString := ""

	m := make(map[byte]int)
	for i, j := 0, 0; i < len(fullString) && j < len(fullString); i++ {
		m[fullString[i]] += 1
		if 1 < m[fullString[i]] {
			nonRepeatingStr := fullString[j:i]
			if 0 == len(maxString) || len(maxString) < len(nonRepeatingStr) {
				maxString = nonRepeatingStr
			}
			j = i
		}
	}
	return maxString
}

func main() {
	fmt.Println(longestNonRepeatSubstr("abcabcbb"))
	fmt.Println(longestNonRepeatSubstr("bbbbb"))
	fmt.Println(longestNonRepeatSubstr("pwwkew"))
}
