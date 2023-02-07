package main

import (
	"fmt"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

/*
Base case:

	if left and right indices are equal, return 1
	if left is greater than right index, return 0

Recursive:

	if s(L) == s(R)
	   return 2 + rec(L+1, R-1)
	else
	   return max (rec(L+1, R), rec(L, R-1))
*/
func recurse(s string, l int, r int) int {
	if l == r {
		return 1
	}
	if l > r {
		return 0
	}

	if s[l] == s[r] {
		return 2 + recurse(s, l+1, r-1)
	} else {
		lRes := recurse(s, l+1, r)
		rRes := recurse(s, l, r-1)
		return Max(lRes, rRes)
	}
}

func main() {

	s := "ttbtctcbt"
	fmt.Println(recurse(s, 0, len(s)-1))
}
