/**

Minimum Window Substring

Given a string and a set of characters return the substring representing the smallest windows containing those characters

Example 1:

Input: "ADOBECODEBANC", "ABC"
Output: "BANC"

Explanation:
Though there are many substrings containing all the characters "BANC" is the shortest.

Example 2:

Input: "HELLO WORLD", "FOO"
Output: ""

Explanation:
The target characters are not contained within the original string, so the output is empty.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func minimumWindowSubstring(fullString string, chars string) string {

	fmt.Printf("%s, %s\n", fullString, chars)

	minString := ""

	for i := 0; i < len(fullString); i++ {

		indices := make([]int, len(chars))
		workingStr := fullString[i:]

		//the working string length must be long enough to accommodate the target substring
		if len(workingStr) < len(chars) {
			break
		}

		for j := 0; j < len(chars); j++ {
			indices[j] = strings.Index(workingStr, string(chars[j]))
			if -1 == indices[j] {
				return minString
			}
		}

		sort.Ints(indices)
		sWindow := workingStr[indices[0] : indices[len(indices)-1]+1]
		fmt.Printf("%03d %-10s (%s)\n", i+1, sWindow, workingStr)

		if 0 < len(sWindow) && (0 == len(minString) || len(sWindow) < len(minString)) {
			minString = sWindow
		}
	}

	return minString
}

func main() {
	//fmt.Println(minimumWindowSubstring("HELLO WORLD", "DLR"))
	fmt.Println(minimumWindowSubstring("ADOBECODEBANC", "ABC"))
	fmt.Println(minimumWindowSubstring("HELLO WORLD", "FOO"))
}
